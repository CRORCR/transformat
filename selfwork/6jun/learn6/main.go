package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"strings"
)

func main() {
	//strDemo()
	md5AndSha()

}

func strDemo() {
	//速度最快
	var buf bytes.Buffer
	buf.WriteString("hello ")
	buf.WriteString("world")
	fmt.Println(buf.String()) //hello world

	//使用join连接
	str := strings.Join([]string{"hello", " world"}, "")
	fmt.Println(str)

	//最low
	str2 := "hello"
	str += "roc"
	fmt.Println(str2)
}

func md5AndSha() {
	str := "helloworld"
	sumMd5 := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	fmt.Println(sumMd5)

	str2 := "helloworld"
	sha := sha1.Sum([]byte(str2))
	shaStr := fmt.Sprintf("%x", sha) //6adfb183a4a2c94a2f92dab5ade762a47889a5a1
	fmt.Println(shaStr)
}
