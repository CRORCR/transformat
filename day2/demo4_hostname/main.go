package main

import (
	"fmt"
	"os"
)

func main() {
	//获取主机名,返回字符串和error信息
	name, err := os.Hostname()
	fmt.Printf("%s %v\n", name, err) //SC-HFJS09241601 <nil>

	//获得环境变量的PATH信息
	val := os.Getenv("PATH")
	fmt.Printf("%s\n", val)
}
