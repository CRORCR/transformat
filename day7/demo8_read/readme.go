package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fileRead()
	fileBufio()
}

func fileBufio() {
	f, err := os.Open("readme.txt")
	if err == nil {
		fmt.Println("打开文件有误")
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取文件遇到异常")
		}
		fmt.Println(line)
	}
}
func fileRead() {
	f, err := os.Open("readme.txt")
	if err == nil {
		fmt.Println("打开文件有误")
	}
	defer f.Close()
	var data [1024]byte
	for {
		n, err := f.Read(data[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取文件遇到异常")
		}
		str := string(data[0:n])
		fmt.Println(str)
	}
}
