package main

import (
	"fmt"

	"github.com/astaxie/beego"
)

func main() {
	//beego.BConfig.Listen.HTTPPort = 9000
	fmt.Println("hello world")
	beego.Run()
}