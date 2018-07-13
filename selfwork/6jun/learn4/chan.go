package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup
var ch = make(chan int, 10)

func main() {
	wait.Add(10)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
			wait.Done()
		}

	}()
	wait.Add(1)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		wait.Done()
	}()

	wait.Wait()
	fmt.Println("exit")
}
