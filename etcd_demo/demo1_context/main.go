package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)
//使用context处理超时

type Result struct {
	r *http.Response
	err error
}
func main() {
	//1.设置超时时间
	ctx,cancel:=context.WithTimeout(context.Background(),2*time.Second)
	defer cancel()

	//2.创建客户端,访问连接 client
	tr:=&http.Transport{}
	client:=&http.Client{Transport:tr}
	c:=make(chan Result,1)

	//创建一个请求连接
	req,err:=http.NewRequest("GET","https://www.baidu.com/",nil)
	if err!=nil{
		fmt.Println("http request failed,err: ",err)
		return
	}
	go func(){
		//访问连接,获得返回response
		resp,err:=client.Do(req)
		//结果放入管道
		c<-Result{r:resp,err:err}
	}()

	//3.轮循,先有错误还是先有数据
	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		fmt.Println("timeout")
	case res:=<-c:
			defer res.r.Body.Close()
			out,_:=ioutil.ReadAll(res.r.Body)
			fmt.Println("server response:%s",out)
	}
	return
}
