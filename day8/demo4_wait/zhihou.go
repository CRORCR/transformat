package main

import (
	"fmt"
	"sync"
	"time"
)

//使用wait第二次优化goroute,当进入协程加1,协程结束减一
//判断什么时候wait(等于0,就退出),代码看起来优雅,逻辑清晰
var wait sync.WaitGroup

func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 3; i++ {
		wait.Add(1)
		go calc()
	}
	wait.Wait()
	end := time.Now().UnixNano()
	fmt.Printf("运行时间%d \n", (end-start)/1000/1000)
}
func calc() {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Microsecond)
	}
	wait.Done()
}
