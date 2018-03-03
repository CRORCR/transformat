package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12/5212/Go"
	format                 = "%f / %d /%s "
)

//数据来源:终端或者字符串都可以,演示如下:
func main() {
	//sprintln内部就是使用了stdout,输出到控制台
	fmt.Println("Please 	enter your full name : ")
	//从控制台读取
	fmt.Scanln(&firstName, &lastName)
	//也是从控制台读取,并格式化,可以指定格式
	fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("HI %s ,%s \n", firstName, lastName) //HI lcq ,李长全
	//从字符串读取,并格式化(解析字符串)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read :", f, i, s)
}
