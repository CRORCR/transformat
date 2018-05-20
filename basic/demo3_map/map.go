package main

import (
	"fmt"
)

//寻找最长不含有重复字符的子串
//解析:
//对于每一个字母X
//1.最后出现的位置找不到,或者在start之前,不需要任何操作
//2.如果最后出现的位置在start和x之间,那就需要更新start位置(最后出现位置+1)
//3.更新最后出现位置,更新最大长度

func main() {
	fmt.Println(getLength("ppp"))
	fmt.Println(getLength("abcbd"))
	fmt.Println(getLength("我是中国人国人"))
}

func getLength(s string) int {
	lastcx := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		//1.如果最后出现的位置在start之后,就更新start位置
		lastI, ok := lastcx[ch] //是否存在,不存在最好了,不用处理
		if ok && lastI >= start {
			start = lastI + 1
		}
		//2.校验length
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		//3.更新map 最后出现的位置
		lastcx[ch] = i
	}
	return maxLength
}
