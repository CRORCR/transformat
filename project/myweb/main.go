package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	_ "transformat/project/myweb/routers"
)

func main() {
	//设置模板
	beego.TestBeegoInit("xxx")
	//设置静态文件
	beego.SetStaticPath("/code", "code")
	//是否显示静态文件
	beego.BConfig.WebConfig.DirectoryIndex = true

	host := beego.AppConfig.String("dbconfig::mysqlhost")
	port, _ := beego.AppConfig.Int("dbconfig::mysqlport")

	//获得int值,如果没有就设置默认值
	redisPort := beego.AppConfig.DefaultInt("dbconfig::redisport", 10000)
	fmt.Println(host, port, redisPort)

	//配置日志
	m := make(map[string]interface{})
	m["filename"] = "./logs/test.log"
	config, _ := json.Marshal(m)

	beego.SetLogger("file", string(config))
	beego.SetLevel(beego.LevelInformational)
	//设置输出的日志带上文件名
	beego.SetLogFuncCall(true)
	beego.Run()
}
