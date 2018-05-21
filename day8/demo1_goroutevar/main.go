package main

import (
	"fmt"
	"time"
)

var exits [3]bool

func main() {
	go calc(0)
	go calc(1)
	go calc(2)
	for {
		if exits[0] == true && exits[1] == true && exits[2] == true {
			break
		}
	}
	fmt.Println("success")
}

func calc(index int) {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond)
	}
	exits[index] = true
}
