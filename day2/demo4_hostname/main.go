package main

import (
	"fmt"
	"os"
)

func main() {
	getUser()
	getPath()
}

//获取主机名
func getUser() {
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Printf("获取主机名有误,错误:%v \n", err)
	}
	fmt.Printf("主机名为:%s \n", hostName)

}

//获取环境变量path
func getPath() {
	path := os.Getenv("PATH")
	fmt.Printf("path:%s \n", path)
}
