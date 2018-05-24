package main

import (
	"fmt"
	"transformat/basic/demo6_func/fibo/fifi"
)

func main() {
	f := fifi.Fibo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
