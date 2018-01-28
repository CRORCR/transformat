package main

import (
	"fmt"
	"strings"
)

func main() {
	//getScan()
	//strOperate()
	char()
}

//获取控制台输入的字符串,并打印出来
func getScan() {
	var a, b string
	fmt.Scanf("%s %s:", &a, &b)
	fmt.Println(a, b)
}

//字符串操作
func strOperate() {
	str := "hello "
	fmt.Printf("str长度:%d \n", len(str)) //6,包括空格

	str2 := str + "world"
	fmt.Printf("str2:%s \n", str2) //hello world

	str3 := "hello,world,你好,世界"
	result := strings.Split(str3, ",") //split返回是字符串数组  字符串操作使用strings
	fmt.Printf("result:%s \n", result) //[hello world 你好 世界]

	str4 := strings.Join(result, ",")
	fmt.Printf("str4:%s \n", str4) //hello,world,你好,世界

	isContain := strings.Contains(str4, "世界")
	fmt.Printf("str4是否包含 世界 关键字,%t \n", isContain) //%t bool格式化
	fmt.Printf("str4是否包含 世界 关键字,%v \n", isContain) //%v万能格式化

	index := strings.Index(str4, "hello")
	fmt.Printf("str4是否存在hello关键字,如果存在,输出角标%d \n", index) //0

	isPrefix := strings.HasPrefix(str4, "hello")
	fmt.Printf("是否以关键字开始:%v \n", isPrefix)

	isSuffix := strings.HasSuffix(str4, "世界")
	fmt.Printf("是否以关键字结束:%v \n", isSuffix)
}

func char() {
	str := "hello"
	for index, v := range str {
		fmt.Printf("角标:%d, 字符:%c \n", index, v)
	}
}
