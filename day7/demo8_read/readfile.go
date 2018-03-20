package main

import (
	"fmt"
	"io/ioutil"
)

//一次读取一个文件 一般用来读取配置文件
func main() {
	data, err := ioutil.ReadFile("readme.txt")
	if err != nil {
		fmt.Println("读取文件有误!!")
	}
	fmt.Println(string(data))
}
