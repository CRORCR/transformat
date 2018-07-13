package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := os.Open(".\\hello.txt")
	defer file.Close()
	//第一种:裸读取
	buff := make([]byte, 10)
	n, _ := file.Read(buff)
	fmt.Println(buff[:n])
	//第二种 buffio
	reader := bufio.NewReader(file)
	reader.ReadString('\n')
	//第三种 读取所有
	bytes, _ := ioutil.ReadAll(file)
	fmt.Println(string(bytes))
}
