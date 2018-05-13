package main

import (
	"fmt"
	"regexp"
)

const text = `my email is 112@qq.com
email2 is quantin@163.com
email3 ting@163.com.cn
`

func main() {
	//regexpBasic()
	//regexOne()
	//findAllStirng()

	//获取@之前和之后,以及域名点之前和之后的字符,匹配项全部拆开输出
	//例如: 112@qq.com.cn   112 qq.com cn
	compile := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	//findAll 匹配所有 参数-1表示所有
	match := compile.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	//[112@qq.com.cn 112 qq.com cn]  全部能匹配上
}
func findAllStirng() {
	compile := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	//findAll 匹配所有 参数-1表示所有
	match := compile.FindAllString(text, -1)
	fmt.Println(match)
	//112@qq.com quantin@163.com ting@163.com
}
func regexOne() {
	//这里需要把.转义 +:1--多个  *:0--多个
	compile := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]+`)
	match := compile.FindString(text)
	//这种方式只能匹配一个
	fmt.Println(match) //112@qq.com
}
func regexpBasic() {
	//获得正则表达式匹配器和error
	//一般匹配用户输入信息,才需要匹配错误
	// 自己写的正则知道是什么格式,直接使用myst匹配,如果匹配不上直接panic
	//compile, err := regexp.Compile(text)
	re := regexp.MustCompile("112@qq.com")
	match := re.FindString(text) //从源文件匹配符合校验器规则的字符串
	fmt.Println(match)           //112@qq.com
}
