package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	log "k8s.io/klog/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
docs: https://github.com/etcd-io/etcd/blob/main/tests/integration/clientv3/concurrency/example_election_test.go

几个方法:
1. 创建一个session: s1, err := concurrency.NewSession(cli)
2. 创建一个选举对象 e1 := concurrency.NewElection(s1, "/my-election/")
3. 尝试进行选举, 如果成功，执行xxx; if err := e1.Campaign(context.Background(), "e1"); err == nil { xxx}
4. 手动放弃主，运行其他进程选为主: e.Resign(context.TODO())
*/

func run_or_die(c *clientv3.Client, key_prefix string, f func(ctx context.Context) error, ctx context.Context) error {
	hostname, _ := os.Hostname()
	client_name := fmt.Sprintf("%s-%d", hostname, time.Now().Nanosecond())

	// 新建一个session
	s, err := concurrency.NewSession(c, concurrency.WithTTL(3))
	if err != nil {
		fmt.Println(err)
	}
	defer s.Close()

	// 新建一个选主对象
	e := concurrency.NewElection(s, key_prefix)
	done := make(chan bool)
	ctx2, cancel := context.WithCancel(ctx)
	defer cancel()

	log.Infof("try get leader %s", client_name)
	if err = e.Campaign(ctx2, client_name); err == nil {
		log.Infof("get leader %s", client_name)
		go func(ctx3 context.Context, out chan<- bool) {
			err2 := f(ctx3) // 正常情况，会一直运行程序在这里

			select {
			case <-ctx3.Done():
				log.Infof("leader task cancel")
			case out <- true:
				log.Infof("leader task finish with err: %v", err2)
			}

		}(ctx2, done)
	} else {
		log.Infof("get leader err: %s", err)
	}

	for {
		select {
		case <-s.Done(): //获取 leader 运行程序后, 需要检测: 在无法连接etcd 后，自动退出程序，避免多个task 运行
			log.Infof("run or dir finished with session.done()=true.")
			return nil
		case <-ctx.Done(): // 主程序希望退出时（用户取消或kill进程等）, 退出任务
			log.Infof("run or dir finished with done=true.")
			return nil
		case <-done: // 任务程序结束或异常退出, 退出任务
			log.Infof("run or dir finished with done_task=true.")
			e.Resign(context.TODO())
			return nil
		}
	}
}

func wait_signal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	log.Infof("开始监听退出信号")
	sig := <-sigs
	log.Infof("收到信号 %v", sig)
}

func demo_run_or_die(c *clientv3.Client) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		//TODO 监听信号等进行 cancel
		wait_signal()
		cancel()

	}()

	run_or_die(c, "/election-demo", demo_task, ctx)

	return nil
}

func demo_task(ctx context.Context) error {
	log.Infof("start task")
	select {
	case <-time.After(60 * time.Second):
		log.Infof("finish task")
	case <-ctx.Done():
		log.Infof("task cancel")
	}

	return nil
}
