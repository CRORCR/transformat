package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type LogConfig struct {
	Topic    string `json:"topic"`
	LogPath  string `json:"log_path"`
	Service  string `json:"service"`
	SendRate int    `json:"send_rate"`
}

var client *clientv3.Client
var logConfChan chan string
var waitGroup sync.WaitGroup

func initEtcd() (err error) {
	addr := appConfig.EtcdAddr
	timeStr := time.Duration(appConfig.etcdTimeout)
	duration := timeStr * time.Second

	keyFmt := appConfig.etcdWatchKeyFmt
	var keys []string
	for _, ip := range ipAarrays {
		keys = append(keys, fmt.Sprintf(keyFmt, ip))
	}

	client, err := clientv3.New(clientv3.Config{Endpoints: addr, DialTimeout: duration})
	if err != nil {
		return
	}
	waitGroup.Add(1)
	for _, key := range keys {
		ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)

		response, err := client.Get(ctx, key)
		cancelFunc()
		if err != nil {
			continue
		}
		for _, v := range response.Kvs {
			logConfChan <- string(v.Value)
		}
	}
	go watchEtcd(keys)
	return
}

func watchEtcd(keys []string) {
	//需要监听那些端口?
	var watchChans []clientv3.WatchChan
	for _, key := range keys {
		rch := client.Watch(context.Background(), key)
		watchChans = append(watchChans, rch)
	}
	//循环去监听
	for {
		for _, watchC := range watchChans {
			select {
			case wresp := <-watchC:
				for _, v := range wresp.Events {
					logConfChan <- string(v.Kv.Value)
				}
			default:

			}
		}
		time.Sleep(time.Second)
	}
	waitGroup.Done()
}

func GetLogConfChan() chan string {
	return logConfChan
}
