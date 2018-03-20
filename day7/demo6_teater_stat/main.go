package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader:=bufio.NewReader(os.Stdin)
	line,_:=reader.ReadString('\n')
	stat(line)
}

func stat(str string){
	var en_count int
	var sp_count int
	var num_count int
	var other_count int
	utf8Arr:=[]rune(str)//字符串转换成数组
	for _,v :=range utf8Arr{
		if v>='a' && v<='z' || v>='A' && v<='Z'{//如果是英文字符
			en_count++
			continue
		}
		if v==' '{
			sp_count++
			continue
		}
		if v>='0' && v<='9'{
			num_count++
			continue
		}
		other_count++
	}
	fmt.Printf("英文=%d 空格=%d 数字=%d 其他字符=%d",en_count,sp_count,num_count,other_count)
}