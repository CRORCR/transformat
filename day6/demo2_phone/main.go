package main

import (
	"fmt"
)

func main() {

}

func demo1() {
	//1.实例化手机
	phone := &Phone{
		//初始化map
		PayMap: make(map[string]Pay, 16),
	}
	//2.开通微信支付
	phone.OpenWeChatPay()
	//3.微信支付
	err := phone.PayMoney("wechat_pay", 20.32)
	if err != nil {
		fmt.Printf("支付失败，失败原因:%v\n", err)
		//4.微信支付失败,使用支付宝
		fmt.Printf("使用支付宝支付\n")
		err = phone.PayMoney("ali_pay", 20.32)
		if err != nil {
			fmt.Printf("支付失败，失败原因:%v\n", err)
			return
		}
	}
	fmt.Println("支付成功，欢迎再次光临！")
}

//微信支付失败,使用支付宝支付
func demo2() {
	//1.实例化手机
	phone := &Phone{
		//初始化map
		PayMap: make(map[string]Pay, 16),
	}
	//2.开通支付宝支付
	phone.OpenAliPay()
	//3.微信支付
	err := phone.PayMoney("wechat_pay", 20.32)
	if err != nil {
		fmt.Printf("支付失败，失败原因:%v\n", err)
		//4.微信支付失败,使用支付宝
		fmt.Printf("使用支付宝支付\n")
		err = phone.PayMoney("ali_pay", 20.32)
		if err != nil {
			fmt.Printf("支付失败，失败原因:%v\n", err)
			return
		}
	}
	fmt.Println("支付成功，欢迎再次光临！")
}

//优化后
func demo3() {
	//1.实例化手机
	phone := &Phone{
		PayMap: make(map[string]Pay, 16),
	}
	//2.判断微信支付,是否实现了pay接口
	//先创建实例--->转为接口-->接口之间类型判断
	weChat := &WeChatPay{}
	var tmp interface{} = weChat
	_, ok := tmp.(Pay)
	if ok {
		fmt.Println("weChat is implement Pay interface")
		//phone.OpenPay("wechat_pay", weChat)
	}
	phone.OpenPay("ali_pay", &AliPay{})

	err := phone.PayMoney("wechat_pay", 20.32)
	if err != nil {
		fmt.Printf("支付失败，失败原因:%v\n", err)
		fmt.Printf("使用支付宝支付\n")
		err = phone.PayMoney("ali_pay", 20.32)
		if err != nil {
			fmt.Printf("支付失败，失败原因:%v\n", err)
			return
		}
	}

	fmt.Println("支付成功，欢迎再次光临！")
}
