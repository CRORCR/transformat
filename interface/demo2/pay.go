package main

type Pay interface {
	//用户标识 扣钱数量 可能的异常
	Pay(userid int32, money float32) error
}
