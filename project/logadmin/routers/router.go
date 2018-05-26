package routers

import (
	"github.com/astaxie/beego"
	"transformat/project/logadmin/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/app/list", &controllers.LogController{}, "*:ListApp")

	beego.Router("/log/list", &controllers.LogController{}, "*:ListLog")
	beego.Router("/app/create", &controllers.LogController{}, "*:CreateApp")
	beego.Router("/app/submit", &controllers.LogController{}, "*:SubmitApp")

	beego.Router("/etcd/submit", &controllers.LogController{}, "*:SubmitEtcd")
	beego.Router("/", &controllers.LogController{}, "*:ListApp")
}
