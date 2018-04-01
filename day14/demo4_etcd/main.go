package main

import (
	"fmt"
	etcd3 "github.com/coreos/etcd/clientv3"
	"time"
)
//etcd使用
//Endpoints 是etcd端口 etcd是集群模式,可能有多个,所以是数组结构
//
func main() {
	cli,err:=etcd3.New(etcd3.Config{
		Endpoints:[]string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout:5*time.Second,
	})

	if err!=nil{
		fmt.Printf("connection is failed,err:%v \n",err)
		return
	}
	fmt.Println("connection success")
	defer cli.Close()
}
