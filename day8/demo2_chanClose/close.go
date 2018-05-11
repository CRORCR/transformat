package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	//chan关闭之后,不能写,只能读
	close(ch)
	for {
		var b int
		//判断管道是否关闭,如果关闭就退出,否则死循环
		b,ok := <-ch
		if ok == false {
			fmt.Println("chan is close")
			break
		}
		fmt.Println(b)
	}
}
