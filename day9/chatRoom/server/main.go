package server

import (
	"fmt"
	"net"
)

var (
	clientMgr *ClientMgr
)

func main() {
	clientMgr = NewClientMgr(10)
	fmt.Printf("start server ....\n")
	//1.监听端口
	l, err := startServer("192.168.8.200:8080")
	if err != nil {
		fmt.Printf("监听端口有误")
		return
	}
	//开启服务
	err = runServer(l)
	if err != nil {
		fmt.Printf("run server is failed err:%v \n", err)
		return
	}
	fmt.Println("server is exied")
}

func startServer(addr string) (l net.Listener, err error) {
	l, err = net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("listen addr:%s failed, err:%v\n", addr, err)
		return
	}
	return
}
