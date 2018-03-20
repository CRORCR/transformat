package main

import (
	"fmt"
	"regexp"
)

var str string

//从终端获取一行字符串,统计英文,数字,空格和其他的数量
func main() {

	//demo()
	regexpdemo()
	//fmt.Println("Please 	enter string : ")
	//fmt.Scanln(&str)
	////英文 / 数字 /空格 /其他
	//len1, len2, len3, len4 := 0, 0, 0, 0
	//for i := 0; i < len(str); i++ {
	//	s := string(str[i])
	//	fmt.Println(s)
	//	if match, _ := regexp.MatchString("[a-z]", s); match {
	//		len1++
	//	} else if match, _ := regexp.MatchString("[0-9]", s); match {
	//		len2++
	//	} else if match, _ := regexp.MatchString("[ ]", s); match {
	//		len3++
	//	} else {
	//		len4++
	//	}
	//}
	//fmt.Printf("英文字符:%d ; 数字:%d ; 空格: %d ; 其他:%d", len1, len2, len3)
}

func regexpdemo() {
	fmt.Println("Please 	enter string : ")
	fmt.Scanln(&str)
	//英文 / 数字 /空格 /其他
	len1, len2, len3 := 0, 0, 0
	fmt.Println(len(str))
	for i := 0; i < len(str); i++ {
		s := string(str[i])
		matchen, _ := regexp.MatchString("[a-z]", s)
		if matchen {
			len1++
		}
		matchnum, _ := regexp.MatchString("[0-9]", s)
		if matchnum {
			len2++
		}
		matchnil, _ := regexp.MatchString("[ ]", s)
		if matchnil {
			len3++
		}
	}
	fmt.Printf("英文字符:%d ; 数字:%d ; 空格: %d ; 其他:%d", len1, len2, len3)
}

func demo() {
	match, _ := regexp.MatchString("[a-z]", "peach")
	fmt.Print(match)
}
