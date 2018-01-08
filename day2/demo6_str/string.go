package main

import "fmt"

//字符串有两种表示形式,"" 这种比较常见,还有一种`` 这种方式换行也会输出,不需要其他连接符
func main() {

	var str string
	str = "abc\n"

	fmt.Printf("%s\n", str)

	var str2 string
	str2 = `abc


	hello


	\n`
	fmt.Printf("%s\n", str2)
}
