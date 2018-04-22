package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁学习
var rlock sync.RWMutex
var w sync.WaitGroup
var count int
func main() {
	w.Add(1)
	t1:=time.Now().UnixNano()
	go func(){
		for i:=0;i<100000;i++{
			rlock.Lock()
			count++
			rlock.Unlock()
		}
		w.Done()
	}()

	for i:=0;i<10;i++{
		go func(){
			for{
				rlock.RLock()
				fmt.Println(count)
				rlock.RUnlock()
			}
		}()
	}
	w.Wait()
	t2:=time.Now().UnixNano()
	fmt.Println((t2-t1)/1000/1000)
}