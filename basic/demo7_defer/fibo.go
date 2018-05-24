package main

import (
	"bufio"
	"fmt"
	"os"
	"transformat/basic/demo6_func/fibo/fifi"
)

func main() {
	write("fib.txt")
}

func write(fileName string) {
	//file, err := os.Create(fileName)
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fifi.Fibo()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}
