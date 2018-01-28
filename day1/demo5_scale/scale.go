package main

import "fmt"

//%d 10进制
//%b 2进制
//%x 16进制
//%f 浮点型
func main() {
	var str string = "hello world"
	var in int = 100
	var f float32 = 3
	fmt.Printf("字符串输出:[ %s ] \n", str)
	fmt.Printf("二进制输出:[ %b ] \n", in)
	fmt.Printf("十进制输出:[ %d ] \n", in)
	fmt.Printf("十六进制输出:[ %x ] \n", in)
	fmt.Printf("浮点型输出:[ %f ] \n", f)

}
