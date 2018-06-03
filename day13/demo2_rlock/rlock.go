package main

import (
	"fmt"
	"sync"
)

var w sync.WaitGroup
var count int
var lock sync.RWMutex

func main() {
	w.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			lock.Lock()
			count++
			lock.Unlock()
		}
		w.Done()
	}()
	//启动16个
	w.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			lock.RLock()
			fmt.Println(count)
			lock.RUnlock()
		}
		w.Done()
	}()
	w.Wait()
	fmt.Println(count)
}
