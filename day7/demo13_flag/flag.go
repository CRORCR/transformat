package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	var s string
	var i int
	flag.String(&s, "D:\readme.config", "print on line")
	flag.IntVar(&i, "", 8, "print on line int")
	flag.Parse()
}
func main() {
	//demo1()
	demo2()
}

func demo2() {

}
func demo1() {
	for index, value := range os.Args {
		fmt.Println("index%s;value%s", index, value)
	}
}
