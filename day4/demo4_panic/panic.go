package main

import "fmt"
import "time"
import "errors"

func initConfig() (err error) {
	return errors.New("init config failed")
}

//recover抓取异常必须在defer函数中才能被调用
func main() {
	for {
		test()
		time.Sleep(time.Second)
	}
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := initConfig()
	if err != nil {
		panic(err)
	}
	return
}
