package main

import "fmt"

//闭包
func main() {
	a := add()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}

func add() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}
