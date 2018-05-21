package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)  //大小为0,表示没有缓冲区
	ch2 := make(chan int, 0) //大小为0,表示没有缓冲区
	go producer(ch, ch2)
	go consumer(ch, ch2)
	time.Sleep(time.Second * 5)
}

func producer(ch, ch2 chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(2*i - 1)
		ch <- 1
		<-ch2
	}
}

func consumer(ch, ch2 chan int) {
	for i := 1; i <= 10; i++ {
		<-ch
		ch2 <- 1
		fmt.Println(2 * i)
	}
}
