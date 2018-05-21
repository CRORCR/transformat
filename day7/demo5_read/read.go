package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"time"
)

type ReadStr struct{}

func (r *ReadStr) Read(p []byte) (n int, err error) {
	fmt.Printf("len(p)=%d  :\n", len(p))

	source := "abcdefghigklmnopq"
	for i := 0; i < 32; i++ {
		index := rand.Intn(len(source)) //随机生成一个下标
		p[i] = source[index]            //这个p切片已经初始化了,只能修改值,不能append添加
	}
	p[32] = '\n' //加上换行符就可以使用ReadLine读取一行
	return len(p), nil
}

func main() {
	rand.Seed(time.Now().UnixNano()) //随机初始化种子,每次取得随机数都不一样
	var randstr = &ReadStr{}         //randStr实现了reader接口的read方法
	randReader := bufio.NewReader(randstr)
	line, prefix, _ := randReader.ReadLine()
	fmt.Printf("rand:%s prefix:%v", line, prefix)
}
