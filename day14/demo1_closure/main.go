package main

import (
	"fmt"
	"sync"
)
var wait sync.WaitGroup

func main() {
	for i:=0;i<5;i++{
		wait.Add(1)
		go func(a int){
			fmt.Println(a)
			wait.Done()
		}(i)
	}
	wait.Wait()
}
