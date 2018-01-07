package main

import "fmt"

func main() {
	phone := &Phone{payMap: make(map[string]Pay, 16)}
	phone.OpenAliPay()
	err := phone.PayMoney("支付宝", 20.32)
	if err != nil {
		fmt.Printf("支付失败,原因是:%s", err)
		fmt.Println("尝试使用微信支付")
		phone.OpenWeixin()
		err = phone.PayMoney("微信", 20.34)
		if err != nil {
			fmt.Printf("支付失败,原因是:%s", err)
			return
		}
	}
	fmt.Println("支付成功,欢迎下次光临")
}
