package route

import (
	"transformat/seckill/proxy/controller"
	"github.com/astaxie/beego"
)

func init(){
	beego.Router("/sec_info/",&controller.ProxyController{},"*:SecInfo")
	beego.Router("/sec_kill/",&controller.ProxyController{},"*:SecKill")
}
