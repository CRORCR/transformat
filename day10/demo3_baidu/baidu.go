package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//客户端
func main() {
	res, err := http.Get("http://baidu.com/")
	if err != nil {
		fmt.Printf("get err:%v \n", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("get data is failed ,err:%v \n", err)
		return
	}
	fmt.Println(string(data))
}
