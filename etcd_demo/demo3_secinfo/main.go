package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var logconf = `
[
    {
        "topic":"nginx_log",
        "log_path":"D:/opensource/nginx-1.13.10/logs/error.log",
        "service":"account",
        "send_rate":50000
	},
	{
        "topic":"nginx_log",
        "log_path":"D:/opensource/nginx-1.13.10/logs/access.log",
        "service":"nginx",
        "send_rate":50000
	}
]
`
var transconf = `
[
    "nginx_log"
]
`

var product_conf = `
[{
	"product_id":1,
	"start_time":1523698100,
	"end_time": 1523709100,
	"status": 3
}]`

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {//connect failed
		return
	}
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/seckill/product/conf", product_conf)
	cancel()
	if err != nil {//put failed
		return
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/seckill/", clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
