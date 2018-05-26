package routers

import (
	"github.com/astaxie/beego"
	"transformat/project/myweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user", &controllers.UserController{}, "*:UserInfo")
	beego.Router("/test", &controllers.UserController{}, "*:Test")
}
