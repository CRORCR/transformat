package main

type Pay interface {
	//支付接口 谁 多少钱(浮点数字) 可能有异常
	pay(user_id int64, money float32) error
}
