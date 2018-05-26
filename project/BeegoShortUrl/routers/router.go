package routers

import (
	"github.com/astaxie/beego"
	"transformat/project/BeegoShortUrl/controllers"
)

func init() {
	beego.Router("/trans/long2short", &controllers.ShortUrlController{}, "post:Long2Short")
	beego.Router("/trans/short2long", &controllers.ShortUrlController{}, "post:Short2Long")
	beego.Router("/shorturl", &controllers.ShortUrlController{}, "get:ShortUrlList")
	beego.Router("/jump", &controllers.ShortUrlController{}, "get:Jump")
}
