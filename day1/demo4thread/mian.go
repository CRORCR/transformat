package main

import (
	"fmt"
	"time"
)

//启动一个线程使用go关键字
func getSing(a int) {
	for i := 0; i < a; i++ {
		if i%2 == 1 {
			fmt.Println("单数: ", i)
		}
	}
}

func plural(a int) {
	for i := 0; i < a; i++ {
		if i%2 == 0 {
			fmt.Println("复数: ", i)
		}
	}
}

func main() {
	go getSing(10)
	go plural(10)
	time.Sleep(10 * time.Second)
}
