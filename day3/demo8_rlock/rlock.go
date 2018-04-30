package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)
var rwLock sync.RWMutex

func main() {
	testRWLock()
}
func testRWLock() {
	var a map[int]int
	a = make(map[int]int, 5)
	//计数,读取了多少次,采用原子操作
	var count int32
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10

	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwLock.Lock()
			b[8] = rand.Intn(100)
			rwLock.Unlock()
		}(a)
	}

	for i := 0; i < 10; i++ {
		go func(b map[int]int) {
			for {
				rwLock.RLock()
				fmt.Println(a)
				rwLock.RUnlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count))
}

