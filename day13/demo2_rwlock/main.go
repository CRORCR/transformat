package main


import (
	"time"
	"sync"
	"fmt"
)

var rwLock sync.RWMutex
var lock sync.Mutex
var w sync.WaitGroup
var count int

func main() {
	w.Add(1)
	start := time.Now().UnixNano()
	go func() {
		for i := 0; i < 1000;i++ {
			lock.Lock()
			count++
			time.Sleep(3*time.Millisecond)
			lock.Unlock()
		}
		w.Done()
	}()

	for i := 0; i < 16; i++ {
		w.Add(1)
		go func() {
			for i := 0; i < 5000;i++ {
				//rwLock.RLock()
				lock.Lock()
				time.Sleep(1*time.Millisecond)
				lock.Unlock()
				//rwLock.RUnlock()
			}
			w.Done()
		}()
	}
	w.Wait()
	end := time.Now().UnixNano()
	fmt.Println((end - start)/1000/1000)
	//fmt.Println(count)
}