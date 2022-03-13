package main

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	log "k8s.io/klog/v2"
	"time"
)

var ETCD_DEMO_WATCH_KEY = "demo_etcd_watch"

func watch_task(cli *clientv3.Client, done chan bool) {
	log.Infof("start watch key %s \n", ETCD_DEMO_WATCH_KEY)
	rch := cli.Watch(context.Background(), ETCD_DEMO_WATCH_KEY)
	for {
		select {
		case wresp := <-rch:
			for _, ev := range wresp.Events {
				log.Infof("watch %s msg %s %q : %q\n", ETCD_DEMO_WATCH_KEY, ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		case <-done:
			log.Infof("watch key %s finish\n", ETCD_DEMO_WATCH_KEY)
			break
		}
	}

}

func watch_demo(cli *clientv3.Client) {
	done := make(chan bool)

	go watch_task(cli, done)
	time.Sleep(time.Second)
	log.Infof("send watch key %s \n", ETCD_DEMO_WATCH_KEY)
	etcd_put(cli, ETCD_DEMO_WATCH_KEY, "1")
	time.Sleep(time.Second)
	log.Infof("send watch key %s \n", ETCD_DEMO_WATCH_KEY)
	etcd_put(cli, ETCD_DEMO_WATCH_KEY, "2")
	time.Sleep(time.Second)
	log.Infof("send watch key %s \n", ETCD_DEMO_WATCH_KEY)
	etcd_put(cli, ETCD_DEMO_WATCH_KEY, "3")
	done <- true
}
