package main

import (
	"context"
	"fmt"
	etcd3 "github.com/coreos/etcd/clientv3"
	"time"
)
func main() {
	//建立连接
	client,err:=etcd3.New(etcd3.Config{
		Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err!=nil{
		fmt.Println("connection is failed",err)
		return
	}
	//最后 关闭连接
	defer client.Close()

	//设置超时 1秒  如果是Nanosecond 1纳秒,那么返回就会报超时
	ctx,cancel:=context.WithTimeout(context.Background(),time.Second)
	_,err=client.Put(ctx,"one","1")
	if err!=nil{
		fmt.Println("存储one节点错误")
		return
	}
	cancel()

	//获取值
	ctx,cancel=context.WithTimeout(context.Background(),time.Second)
	//可以轮询监控etcd 一秒get一次
	resp,err:=client.Get(ctx,"one")
	if err!=nil{
		fmt.Println("获取one节点错误")
		return
	}

	cancel()
	//键是目录结构,可能有很多子目录, 所以是数组结构 需要range
	for _,v:=range resp.Kvs{
		fmt.Printf("one节点键:%s 值:%s",v.Key,v.Value)
	}
}
