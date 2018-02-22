package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

func main() {
	ch := make(chan string)
	wait.Add(2)
	go test3(ch)
	//go test4(ch)
	go test5(ch)
	wait.Wait()
}

func test3(ch chan<- string) {
	ch <- "hello"
	ch <- "world"
	ch <- "你好"
	ch <- "世界"
	close(ch)
	wait.Done()
}
func test4(ch <-chan string) {
	for {
		s, ok := <-ch
		if !ok { //判断管道关闭
			fmt.Println("ch is closed \n")
			break
		}
		fmt.Printf("value:%s \n", s)
	}
	wait.Done()
}

//test4 --> test5
//for range 就不用谢ok了
func test5(ch <-chan string) {
	for s := range ch {
		fmt.Printf("value:%s \n", s)
	}
	wait.Done()
}
