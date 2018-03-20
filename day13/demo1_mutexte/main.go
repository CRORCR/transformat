package main

import (
	"fmt"
	"sync"
)

var count int
var lock sync.Mutex
var w sync.WaitGroup
func main() {
	w.Add(1)
	go func(){
		for  i:=0;i<100;i++{
			lock.Lock()
			count++
			lock.Unlock()
		}
		w.Done()
	}()

	for  i:=0;i<100;i++{
		lock.Lock()
		count++
		lock.Unlock()
	}
	w.Wait()
	fmt.Println(count)
}
