package main

import (
	"context"
	"fmt"
)
//使用context保存上下文
func main() {
	//withValue可以保存一个变量,一直传递,后面只需要调用value()方法就可以获取
	ctx:=context.WithValue(context.Background(),"trace_id","12345")
	index:=calc(ctx,388,200)
	fmt.Println(index)//588
}

func calc(ctx context.Context, a, b int) int {
	traceId:=ctx.Value("trace_id").(string)//转换成string
	fmt.Printf("trace_id:%v\n", traceId)//trace_id:12345
	return add(ctx, a, b)
}

func add(ctx context.Context, a, b int) int {
	traceId := ctx.Value("trace_id").(string)
	fmt.Printf("trace_id:%v\n", traceId)//trace_id:12345
	return a + b
}