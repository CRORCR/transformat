package main

import (
	"fmt"
	"time"
)

func main() {
	go calc()
	time.Sleep(time.Second)
	fmt.Println("recover is exited")
}
func calc() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var p *int
	*p = 100
}
