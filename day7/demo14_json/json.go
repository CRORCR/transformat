package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

//序列化字段首字母要大写,跨包如果小写不能访问
type Student struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
}

func main() {
	//marshal()
	unmarshal()
}

func unmarshal() {
	file, err := os.OpenFile("D:/hello.dat", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("读取文件有误,err:%v", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("读取所有文件有误,err:%v", err)
		return
	}
	var stus []*Student
	err = json.Unmarshal(data, stus)
	if err != nil {
		fmt.Printf("反序列化失败,err:%v", err)
		return
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("student:%v", stus[i])
	}
}
func marshal() {
	var stus []*Student
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:  fmt.Sprintf("LCQ%d ", i),
			Age:   rand.Intn(10),
			Score: rand.Float32() * 10,
		}
		stus = append(stus, stu)
	}
	data, err := json.Marshal(stus)
	if err != nil {
		fmt.Printf("序列化失败,error:%v", err)
	}

	//写入文件
	file, err := os.OpenFile("D:/hello.dat", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("打开文件失败,err:%v", err)
	}
	defer file.Close()

	n, err := file.Write(data)
	if err != nil {
		fmt.Printf("写入文件失败,err:%v", err)
	}
	fmt.Printf("write %d,success \n", n)
}
