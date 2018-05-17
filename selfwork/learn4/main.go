package main

import (
	"fmt"
	"time"
)

func producer(c chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(2*i - 1)
		c <- 1
		c <- 1
	}
}
func consumer(c chan int) {
	for i := 1; i <= 10; i++ {
		<-c
		<-c
		fmt.Println(2 * i)
	}
}

func main() {
	c := make(chan int)
	go producer(c)
	go consumer(c)

	time.Sleep(3 * time.Second)
}
