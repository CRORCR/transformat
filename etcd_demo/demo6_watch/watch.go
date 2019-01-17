package main

import (
	"context"
	"fmt"
	"time"
	etcd3 "github.com/coreos/etcd/clientv3"
)

func main() {
	cli,err:=etcd3.New(etcd3.Config{
		Endpoints:[]string{"localhost:2379","localhost:22379","localhost:32379"},
		DialTimeout:5*time.Second,
	})
	if err!=nil{
		fmt.Println("connection is failed,err:",err)
		return
	}
	fmt.Println("connection success")
	defer cli.Close()

	//测试,加上这句代码,在两个地方运行,就可以捕捉到存入的数据和类型
	cli.Put(context.Background(),"/logagent/conf/","abcd")

	//watch 监控etcd
	//watch之后,就与etcd建立了长连接,不间断监控
	//当etcd变化之后,就通知 watch客户端
	rch:=cli.Watch(context.Background(),"/logagent/conf/")
	for wresp :=range rch{
		for _,v := range wresp.Events{
			//PUT "/logagent/conf/" : "abcd"
			fmt.Printf("%s %q : %q \n",v.Type,v.Kv.Key,v.Kv.Value)
		}
	}

}







