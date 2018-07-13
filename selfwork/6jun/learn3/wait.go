package main

import (
	"fmt"
	"sync"
)

var wait sync.WaitGroup

//会输出1-10之间,并且是无序的
func main() {
	for i := 0; i < 10; i++ {
		wait.Add(1)
		go test(i)
	}
	wait.Wait()
	fmt.Println("exit")
}

func test(i int) {
	fmt.Println(i)
	wait.Done()
}
