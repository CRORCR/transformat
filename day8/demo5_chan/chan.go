package demo5_chan

import (
	"fmt"
	"time"
)

func main() {
	//testNotBuf()//存入一个值
	//testOneBuf()//存入一个值
	testOneBufErr() //存入两个值
}

//start input  只会输出一个start input
func testNotBuf() {
	var c chan int = make(chan int, 0)
	go func() {
		fmt.Println("start input")
		c <- 10
		fmt.Println("end output")
	}()
	time.Sleep(5 * time.Second) //暂停五秒
}

//start input
//end output
func testOneBuf() {
	var c chan int = make(chan int, 1)
	go func() {
		fmt.Println("start input")
		c <- 10
		fmt.Println("end output")
	}()
	time.Sleep(5 * time.Second)
}

//start input
func testOneBufErr() {
	var c chan int = make(chan int, 1)
	go func() {
		fmt.Println("start input")
		c <- 10
		c <- 10
		fmt.Println("end output")
	}()
	time.Sleep(5 * time.Second)
}
