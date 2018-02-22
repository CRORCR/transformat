package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	calc1()
	calc1()
	calc1()
	end := time.Now().UnixNano()
	fmt.Printf("finished,cost:%d ms \n", (end-start)/1000/1000)
}
func calc1() {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Microsecond)
	}
}
