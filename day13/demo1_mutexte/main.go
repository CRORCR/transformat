package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var lock sync.Mutex
var w sync.WaitGroup
func main() {
	w.Add(1)
	t1:=time.Now().UnixNano()
	go func(){
		for  i:=0;i<100000;i++{
			lock.Lock()
			count++
			lock.Unlock()
		}
		w.Done()
	}()

	for  i:=0;i<100000;i++{
		lock.Lock()
		count++
		lock.Unlock()
	}
	w.Wait()
	fmt.Println(count)
	t2:=time.Now().UnixNano()
	fmt.Println((t2-t1)/1000/1000)
}
