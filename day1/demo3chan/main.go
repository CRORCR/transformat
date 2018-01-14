package main

import "fmt"

//ch := make(chan string, 3)
//创建chan需要设置初始容量,否则报错
func main() {
	ch := make(chan string, 3)
	ch <- "a"
	ch <- "b"

	c := <-ch
	fmt.Printf("%s\n", c)
}
