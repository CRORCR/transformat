package main

import (
	"encoding/json"
	"fmt"
)

//json输出必须是printf 格式化输出
//字段需要加上 ` json:"" `
//使用marshal函数转换 将接口转为json字符串

//使用Unmarshal函数 json转换对象 返回一个error参数
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User{"李长全", 26}
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println("转换有误", err)
	}
	//{"name":"李长全","age":26}
	fmt.Printf("%s \n", data)

	err = json.Unmarshal(data, &u)
	if err != nil {
		fmt.Println("转换对象有误")
	}
	fmt.Println(u) //{李长全 26}
}
