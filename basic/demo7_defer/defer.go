package main

import "fmt"

//3   2   1   panic
func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error")
}
