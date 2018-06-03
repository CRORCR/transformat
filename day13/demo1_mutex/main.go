package main

import (
	"fmt"
	"sync"
	"time"
)

//互斥锁
var lock sync.Mutex
var count int
var w sync.WaitGroup

func main() {
	start := time.Now().UnixNano()
	w.Add(2)
	go demo1()
	go demo2()
	w.Wait()
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1000 / 1000)
	fmt.Println(count)
}

func demo1() {
	for i := 0; i < 1000000; i++ {
		lock.Lock()
		count++
		lock.Unlock()
	}
	w.Done()
}

func demo2() {
	for i := 0; i < 1000000; i++ {
		lock.Lock()
		count++
		lock.Unlock()
	}
	w.Done()
}
