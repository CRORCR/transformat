package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	copy("d:/readmebuf.txt", "d:/readme.txt")
}

func copy(a, b string) {
	f, err := os.Open(a)
	if err != nil {
		fmt.Println("读取文件错误", err)
	}
	defer f.Close()

	file, err := os.OpenFile(b, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("写入文件有误", err)
	}
	io.Copy(file, f)

}

func writer() {

}
