package main

import (
	"fmt"
	"time"
)

var exist [3]bool

func main() {
	start := time.Now().UnixNano()
	go calc(0)
	go calc(1)
	go calc(2)
	for {
		if exist[0] == true && exist[1] == true && exist[2] == true {
			break
		}
		time.Sleep(time.Microsecond)
	}
	end := time.Now().UnixNano()
	fmt.Printf("运行时间%d \n", (end-start)/1000/1000)
}
func calc(index int) {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Microsecond)
	}
	exist[index] = true
}
