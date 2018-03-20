package main

import (
	"fmt"
	"sync"
)

var rwlock sync.RWMutex
var w sync.WaitGroup
var count int

func main() {
	w.Add(1)
	go func(){
		for  i:=0;i<100;i++{
			rwlock.Lock()
			count++
			rwlock.Unlock()
		}
		w.Done()
	}()

	for i:=0;i<10;i++{
		go func(){
			for{
				rwlock.RLock()
				fmt.Println(count)
				rwlock.RUnlock()
			}

		}()
	}
	w.Wait()
}
