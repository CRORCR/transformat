package main

import (
	"context"
	"fmt"
)

//使用context传递上下文
//上下文是树状结构,前一个上下文作为参数传递就可以了
//如果不传递,默认是新的上下文
func main() {
	//存值
	ctx:=context.WithValue(context.Background(),"test","123")
	ctx=context.WithValue(context.Background(),"session","session_value")

	//ctx是树状结构,可以继承,前一个ctx作为往下传就可以了
	ctx=context.WithValue(ctx,"session","session_value")
	process(ctx)
}

func process(ctx context.Context){
	//取值,返回值是interface,强转成string,如果不存在,给定默认值
	res,ok:=ctx.Value("test").(string)
	if !ok{
		res="hello"
	}
	fmt.Println(res)

	res=ctx.Value("session").(string)
	//session值 = session_value
	fmt.Printf("session值 = %v",res)
}
