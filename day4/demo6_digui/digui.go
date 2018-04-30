package main

import "fmt"

func main() {
	factor()
	recusive(1)
}

func factor(){
	sum:=prin(5)
	fmt.Println(sum)
}

//1-5相加
func prin(n int)(sum int){
	if n==1{
		return 1
	}else{
		return n+prin(n-1)
	}
}

//输出十遍hello world
func recusive(n int){
	fmt.Println("hello world")
	if n>=10{
		return
	}
	recusive(n+1)
}
