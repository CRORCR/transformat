package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Result struct{
	r *http.Response
	err error
}
func main() {
	//1.设置定时器 两秒没有返回就抛出异常
	ctx,cancel:=context.WithTimeout(context.Background(),2*time.Second)
	//方法执行,最后执行defer 把这个定时器取消
	defer cancel()

	//2.生成client
	tr:=&http.Transport{}
	client:=&http.Client{Transport:tr}
	//3.初始化管道,存储结果
	c:=make(chan Result,1)

	//4.创建请求
	req,err:=http.NewRequest("GET","http://baidu.com",nil)
	if err!=nil{
		fmt.Println("http request failed,err:",err)
		return
	}
	//5.开启goroute
	go func(){
		//发送请求,不知道什么时候返回
		resp,err:=client.Do(req)
		pack:=Result{r:resp,err:err}
		c<-pack
	}()

	//超时先结束 还是 请求先结束
	select {
		case <-ctx.Done():
			//如果超时就取消当前执行的http请求
			tr.CancelRequest(req)
			er:=<-c
			fmt.Println("timeout!!,err:",er.err)
		case res:=<-c:
			//返回结果,关闭连接
			defer res.r.Body.Close()
			out,_:=ioutil.ReadAll(res.r.Body)
			fmt.Printf("server response:%s",out)
	}
}














