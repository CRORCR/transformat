package main

import (
	"fmt"
	"net/http"
)

var url = []string{
	"http://www.baidu.com",
	"http://taobao.com",
}

func main() {
	for _, v := range url {
		res, err := http.Head(v)
		if err != nil {
			fmt.Printf("head is failed err:%v \n", err)
			continue
		}
		fmt.Printf("head success status:%v \n", res.Status)
	}
}
