package main

import (
	"fmt"
	"sync"
	"time"
)

var wait sync.WaitGroup

func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 3; i++ {
		wait.Add(1)
		go calc()

	}
	wait.Wait()
	end := time.Now().UnixNano()
	fmt.Printf("所需时间%d", (end-start)/1000/1000)
}
func calc() {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Microsecond)
	}
	wait.Done()
}
