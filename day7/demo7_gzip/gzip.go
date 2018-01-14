package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("D:\readme.gz")
	if err != nil {
		fmt.Println("打开文件有误:", err)
		return
	}
	defer f.Close()

	reader, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println("读取压缩文件有误", err)
	}
	bufferread := bufio.NewReader(reader)
	for {
		line, err := bufferread.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("读取文件有误", err)
			return
		}
		fmt.Printf("%s \n", line)
	}
}
