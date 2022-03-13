package main

import (
	"context"
	"errors"
	"flag"
	clientv3 "go.etcd.io/etcd/client/v3"
	logs "k8s.io/component-base/logs"
	log "k8s.io/klog/v2"
	"time"
)

/*
docs: https://github.com/etcd-io/etcd/blob/main/tests/integration/clientv3/examples/example_kv_test.go#L185

*/

func etcd_put(cli *clientv3.Client, key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	resp, err := cli.Put(ctx, key, value)
	defer cancel()
	log.Infof("put %s : %+v", key, resp)
	return err
}

func etcd_get(cli *clientv3.Client, key string) (value string, error error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	resp, err := cli.Get(ctx, key)
	defer cancel()
	if err != nil {
		return "", err
	}
	log.Infof("get %s : %+v", key, resp)
	if len(resp.Kvs) != 1 {
		return "", errors.New("many value,try again")
	}
	value = string(resp.Kvs[0].Value)
	return value, nil
}

func etcd_delete(cli *clientv3.Client, key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	resp, err := cli.Delete(ctx, key)
	defer cancel()
	log.Infof("delete %s : %+v", key, resp)
	return err
}

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()
	log.InitFlags(flag.CommandLine)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
	}
	defer cli.Close()

	etcd_put(cli, "test", "2")
	value, _ := etcd_get(cli, "test")
	log.Infof("test: %s", value)
	value2, _ := etcd_get(cli, "test")
	log.Infof("test2: %s", value2)

	err = etcd_delete(cli, "test")
	log.Infof("delete test: %s", err)

	watch_demo(cli)

	demo_run_or_die(cli)
}
