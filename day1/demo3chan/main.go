package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "a"
	ch <- "b"

	c := <-ch
	fmt.Printf("%s\n", c)
}
