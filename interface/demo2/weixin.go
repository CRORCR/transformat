package main

import "fmt"

type WeChatPay struct {
}

func (w *WeChatPay) Pay(userid int32, money float32) error {
	fmt.Println("1.连接到微信支付服务器")
	fmt.Println("2.找到对应用户")
	fmt.Println("3.检查余额")
	fmt.Println("4.扣钱")
	fmt.Println("5.返回是否成功")
	return nil
}
