package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int = make(chan int)
	go queryDB(ch)
	t := time.NewTicker(time.Second)
	//随机扫描,那个值存在,就执行那个go route
	select {
	case v := <-ch:
		fmt.Printf("result: %v", v)
	case <-t.C:
		fmt.Printf("timeOut")
	}
}

func queryDB(ch chan int) {
	//time.Sleep(time.Second)  加上这行就会抛出timeout
	ch <- 100
}
