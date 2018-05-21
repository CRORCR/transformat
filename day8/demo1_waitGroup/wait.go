package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wait sync.WaitGroup
	for i := 0; i < 3; i++ {
		wait.Add(1)
		go calc(&wait)
	}
	wait.Wait()
	fmt.Println("success")
}

func calc(wait *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond)
	}
	wait.Done()
}
