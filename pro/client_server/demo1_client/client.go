package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//客户端处理:
//建立与服务端链接
//收发消息
//关闭连接
func main() {
	//Dial 会话
	//conn, err := net.Dial("tcp", "localhost:8889")
	conn, err := net.Dial("tcp", "192.168.14.200:18080")
	if err != nil {
		fmt.Println("dialing is fail,err:", err.Error())
		return
	}
	defer conn.Close()
	inputRead := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputRead.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimmedInput))
		if err != nil {
			return
		}
	}
}
