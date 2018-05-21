package main

import (
	"fmt"
)

type Phone struct {
	PayMap map[string]Pay
}

// 优化前  OpenWeChatPay() OpenWeChatPay()
func (p *Phone) OpenWeChatPay() {
	weChatPay := &WeChatPay{}
	p.PayMap["wechat_pay"] = weChatPay
}

func (p *Phone) OpenWeChatPay() {
	aliPay := &AliPay{}
	p.PayMap["ali_pay"] = aliPay
}

//优化后 OpenPay()
func (p *Phone) OpenPay(name string, pay Pay) {
	p.PayMap[name] = pay
}

func (p *Phone) PayMoney(name string, money float32) (err error) {
	pay, ok := p.PayMap[name]
	if !ok {
		err = fmt.Errorf("不支持[%s]支付方式", name)
		return
	}

	err = pay.pay(1023, money)
	return
}
