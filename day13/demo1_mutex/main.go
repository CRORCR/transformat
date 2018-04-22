package main

import (
	"fmt"
	"sync"
)


//互斥锁
var lock sync.Mutex
var count int
var w sync.WaitGroup

func main() {
	w.Add(1)
	go demo1()
	go demo2()
	w.Wait()
	fmt.Println(count)
}

func demo1(){
	for i:=0;i<100;i++{
		lock.Lock()
		count++
		lock.Unlock()
	}
	w.Done()
}

func demo2(){
	for i:=0;i<100;i++{
		lock.Lock()
		count++
		lock.Unlock()
	}
}


