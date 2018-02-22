package main

import (
	"encoding/json"
	"fmt"
)

//make用来分配map slice channel类型的内存
//new 用来分配值类型的内存(基本类型,数组,结构体)

//tag:就是标记
//结构体的字段,写是原信息,比如json

//json输出必须是printf 格式化输出
//使用Unmarshal函数 json转换对象 返回一个error参数

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	//序列化
	s := &Student{"lcq", 100}
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Printf("json marshal failed,err:%v \n", err)
		return
	}
	fmt.Printf("%s \n", data) //{"name":"lcq","age":100}

	//反序列化
	var s2 Student
	err = json.Unmarshal(data, &s2)
	if err != nil {
		fmt.Printf("json marshal failed,err:%v \n", err)
		return
	}
	fmt.Println(s2) //{lcq 100}
}
