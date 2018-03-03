package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//缓冲区读写 控制台输入信息
	reads := bufio.NewReader(os.Stdin)
	//按行读取  分割符号是字符,单引号  双引号是字符串
	line, err := reads.ReadString('\n')
	if err != nil {
		fmt.Println("read error: ", err)
	}
	fmt.Println("out", line)
}
