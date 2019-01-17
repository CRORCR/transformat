package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Printf("err is dial err:%v", err)
		return
	}
	defer conn.Close()

	msg := "GET / HTTP/1.1\r\n"
	msg += "Host:www.baidu.com\r\n"
	msg += "Connection:close\r\n"
	msg += "\r\n\r\n"

	_, err = io.WriteString(conn, msg)
	if err != nil {
		fmt.Printf("write is failed err:%v \n", err)
		return
	}
	buf := make([]byte, 1024)
	for {
		count, err := conn.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:count]))
	}
}
