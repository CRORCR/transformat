package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var rwLock sync.RWMutex

func main() {
	testRWLock()
}
func testRWLock() {
	var a map[int]int
	a = make(map[int]int, 5)
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
			}
		}(a)
	}
	time.Sleep(time.Second * 3)
}

