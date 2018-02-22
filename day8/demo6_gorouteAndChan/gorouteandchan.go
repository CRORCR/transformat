package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go test(ch)
	go test2(ch)
	time.Sleep(5 * time.Second)
}
func test(ch chan string) {
	ch <- "hello"
	ch <- "world"
	ch <- "你好"
	ch <- "世界"
}

func test2(ch chan string) {
	var s string
	for {
		s = <-ch
		fmt.Println(s)
	}
}
