package main

import "fmt"

type Phone struct {
	payMap map[string]Pay
}

func (p *Phone) OpenWeixin() {
	weixin := &WeChatPay{}
	p.payMap["微信"] = weixin
}

func (p *Phone) OpenAliPay() {
	ali := &Alipay{}
	p.payMap["支付宝"] = ali
}

func (p *Phone) PayMoney(name string, money float32) (err error) {
	pay, ok := p.payMap[name]
	if !ok {
		err = fmt.Errorf("不支持[%s]方式支付", name)
		return
	}
	err = pay.Pay(1010, money)
	return
}
