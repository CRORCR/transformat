package main

import (
	"fmt"
	"time"
)

//定时器,每隔一秒执行一次
func main() {
	t := time.NewTicker(time.Second)
	for v := range t.C {
		fmt.Println("hello world", v)
	}
}
