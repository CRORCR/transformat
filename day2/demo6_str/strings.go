package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//var url, path string
	//
	//fmt.Scanf("%s%s", &url, &path)
	//urlProcess(url)
	//pathProcess(path)
	strStudy()
	//out()
}

//判断一个url是否以http://开头，如果不是，则加上http://。
func urlProcess(url string) {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}
	fmt.Println(result)
}

//判断一个路径是否以“/”结尾，如果不是，则加上/
func pathProcess(path string) {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	fmt.Println(result)
}

func strStudy() {
	//fmt.Println(strings.TrimSpace(" sksk ")) //-->sksk 中间不管
	//fmt.Println(strings.Trim("abbacba", "ab"))//-->c ab任意组合都会去除 慎用
	//fmt.Println(strings.TrimLeft("b c ","b"))//--> c 去除串首指定字符
	//fmt.Println(strings.TrimRight(" a b c","c"))//-->a b 去除串尾指定字符
	//fmt.Println(strings.Fields("abc cde edk"))//-->[abc cde edk]返回str空格分隔的所有子串的slice
	//fmt.Println(strings.Split("abc,cde,edk", ","))//-->[abc cde edk]
	//fmt.Println(strings.Join([]string{"abc", "cde", "edk"}, ","))//-->abc,cde,edk
	//fmt.Println(strings.Replace("str", "s","3",0))//-->str 替换多少次
	fmt.Println(strings.Repeat("abcdefabcdef", 2)) //重复返回多少次

	//str2 := "hello" + "world"
	//fmt.Printf("str2:%s \n", str2) //hello world
	//isContain := strings.Contains(str2, "世界")
	//fmt.Printf("str2是否包含 hello 关键字,%t \n", isContain) //%t bool格式化
	//fmt.Printf("str2是否包含 hello 关键字,%v \n", isContain) //%v万能格式化
	//strings.ToUpper(str string)string：转为大写
	//strings.ToLower(str string)string：转为小写
	//strings.Count(str string, substr string)int：字符串计数
	//strings.Repeat(str string, count int)string：重复count次str

	//strings.Itoa(i int)：把一个整数i转成字符串
	//strings.Atoi(str string)(int, error)：把一个字符串转成整数
}

//从终端读取字符串,转为int类型,如果出错就返回
func conv() {
	var str string
	fmt.Scanf("%s", &str)

	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("convert failed, err:", err)
		return
	}
	fmt.Println(number)
}

//字符串有两种表示形式,"" 这种比较常见,还有一种`` 这种方式换行也会输出,不需要其他连接符
func out() {

	var str string
	str = "abc\n"

	fmt.Printf("%s\n", str)

	var str2 string
	str2 = `abc


	hello


	\n`
	fmt.Printf("%s\n", str2)
}

//string底层是数组,长度不可变
