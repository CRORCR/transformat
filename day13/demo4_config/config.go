package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//写一个解析器,解析配置文件

type Config struct{
	fileName string
	data map[string]string
}

func Init(filename string)(config *Config,err error){
	config=&Config{
		fileName:filename,
		data:make(map[string]string,16),
	}
	return
}
func ( conf Config)Parse()(err error) {
	//Open 默认只读方式
	file,err:=os.Open(conf.fileName)
	if err!=nil{
		fmt.Println("读取文件失败")
		return
		var lineNo int
		reader:=bufio.NewReader(file)
		for{
			line,err:=reader.ReadString('\n')
			if err==io.EOF{
				break
			}
			if err!=nil{
				return
			}
			lineNo++
			arr:=strings.Split(line,"=")
			if len(arr)==0{
				continue
			}

			}
	}
	return
}
