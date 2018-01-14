package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//test()
	testBuff()
}

func test() {
	file, err := os.OpenFile("d:/readme.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("写入文件有误", err)
	}

	for i := 0; i < 10; i++ {
		file.WriteString(fmt.Sprintf("hello %d \n", i))
	}
}

func testBuff() {
	file, err := os.OpenFile("d:/readmebuf.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("写入文件有误", err)
	}

	out := bufio.NewWriter(file)
	str := ""
	for i := 0; i < 10; i++ {
		str += fmt.Sprintf("hello %d \n", i)
		out.WriteString(str)
	}
	out.Flush()
}
