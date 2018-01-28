package main

import (
	"fmt"
	"time"
)

//ch := make(chan string, 3)
//创建chan需要设置初始容量,否则报错

func main() {
	ch := make(chan int, 3)
	go one(ch)
	go two(ch)
	time.Sleep(1 * time.Second)
}
func one(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
}

func two(ch chan int) {
	//两种方式读取chan值,循环打印控制台
	for c := range ch {
		fmt.Println(c) //1  2  3
	}
	/*for {
		c := <-ch
		fmt.Println(c)//1  2  3
	}*/
}
