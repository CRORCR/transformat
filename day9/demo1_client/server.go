package main

import (
	"fmt"
	"net"
)


//服务端处理:
//监听端口
//接收链接
//创建goroute处理
func main() {
	fmt.Println("client start...")
	listen, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("listen is fail,err: ", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("connection is fail,err: ", err)
			continue
		}
		go proess(conn)
	}
}

func proess(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read is fail ,err: ", err)
			return
		}
		fmt.Println("read:", string(buf))
	}
}


func demo(){

}