package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wait sync.WaitGroup

func main() {
	var ch chan string = make(chan string)
	var exitChan chan bool = make(chan bool, 1)
	var sinalChan chan os.Signal = make(chan os.Signal, 1)
	wait.Add(2)
	//监听指定信号   这个值报错,需要设置两个环境变量 GOARCH=amd64   GOOS=linux
	//然后重启idea ,发现还是报错,但是没有关系,不会影响编译,编译完,放在linux运行
	signal.Notify(sinalChan, syscall.SIGUSR2)

	go test(ch, exitChan)
	go test1(ch)

	<-sinalChan
	exitChan <- true
	wait.Wait()
}

func test(ch chan string, exitChan chan bool) {
	var i int
	var exit bool
	for {
		str := fmt.Sprintf("hello %d", i)
		select {
		case ch <- str:
		case exit = <-exitChan:
		}
		if exit {
			fmt.Printf("user notify product exited \n")
			break
		}
	}
	close(ch)
	wait.Done()
}

func test1(ch chan string) {
	for {
		str, ok := <-ch
		if !ok {
			fmt.Printf("ch is closed \n")
			break
		}
		fmt.Sprintf("value:%s", str)
	}
	wait.Done()
}
