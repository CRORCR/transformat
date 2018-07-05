package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	//size is 26MB
	pathName := "/Users/mfw/Desktop/shakespeare.json"
	start := time.Now()
	readCommon(pathName)
	timeCommon := time.Now()
	log.Printf("read common cost time %v\n", timeCommon.Sub(start))

	readBufio()
	timeBufio := time.Now()
	log.Printf("read bufio cost time %v\n", timeBufio.Sub(timeCommon))

	readIOUtil(pathName)
	timeIOUtil := time.Now()
	log.Printf("read ioutil cost time %v\n", timeIOUtil.Sub(timeBufio))
}

func send(date []byte)(err error){

	file, err := os.OpenFile("lcq.txt", os.O_WRONLY|os.O_CREATE, 0664)
	defer file.Close()
	if err!=nil{
		return
	}
	file.Write(date)
	return
}

//golang读取文件 通过原生态 io 包中的 read 方法进行读取
func readCommon(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		readNum, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == readNum {
			break
		}
	}
}

//通过 io/ioutil 包提供的 read 方法进行读取
func readIOUtil(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = ioutil.ReadAll(file)
}

//总结: 普通读取 和readAll 运行时间相差不大，通过查看源代码发现 io/ioutil 包其实就是封装了 io 包中的方法，故他们两没有性能优先之说；
//性能最好的是通过带有缓冲的 io 流来读取文件。

//通过 bufio 包提供的 read 方法进行读取
func readBufio()(data []byte,rr error){
	data=make([]byte,1024)
	file, err := os.Open("lcq.txt")
	defer file.Close()
	if err!=nil{
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for{
		s, e := reader.ReadString('\n')
		if e==io.EOF{
			break
		}
		by := []byte(s)

		data=append(data,by...)
	}
	return
}
