package main

import (
	"fmt"
	"net"
)

//客户端
//1.监听端口,收发数据
//2.开启goroute
//3.关闭连接

func main() {

	fmt.Println("start server ... ")

	l, err := startServer("0.0.0.0:8080")
	if err != nil {
		fmt.Printf("start server is fail,err:%v \n", err)
	}
	err = runServer(l)
	if err != nil {
		fmt.Printf("run server is fail,err%v \n", err)
	}
	fmt.Println("server is exied")
}

func startServer(address string) (l net.Listener, err error) {
	l, err = net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("listen addr:%s,failed,err:%v \n", address, err)
		return
	}
	return
}
