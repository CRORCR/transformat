package controller



import (
	"fmt"
	"strconv"
	"time"
	"transformat/seckill/proxy/model"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)
//controller定义状态的结构
type ActivityStatus struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int
}

type ProxyController struct{
	beego.Controller
}

//定义错误码,返回map
func (p *ProxyController)back(errno int,message string ,data interface{})(m map[string]interface{}){
	m=make(map[string]interface{},16)
	m["errno"]=errno
	m["message"]=message
	m["data"]=data
	return
}
func (p *ProxyController)SecInfo(){
	logs.Debug("secKill")
	//1.每个分支都需要设置json和server,所以抽取成defer
	m:=make(map[string]interface{},16)
	defer func(){
	p.Data["json"]=m
	p.ServeJSON()
	}()
	//2.获取商品id,如果出错就返回自定义的错误json数据
	product_id,err:=p.GetInt("product_id",0)
	if err!=nil || product_id==0{
		m = p.back(1001, fmt.Sprintf("invalid product_id:%d", product_id), nil)
		logs.Error("invalid product id:%d", product_id)
		return
	}
	//3.根据id获得商品数据
	data,err:=model.SecInfo(product_id)
	if err != nil {
		m = p.back(1002, fmt.Sprintf("service busy, product_id:%d", product_id), nil)
		logs.Error("service busy,product id:%d, err:%v", product_id, err)
		return
	}
	//判断状态
	var activityStatus ActivityStatus
	now:=time.Now().UnixNano()
	//1.当前时间是否在商品活动时间内
	if now>=data.StartTime && now<data.EndTime{
		//1.1.判断状态是不是卖完(常量3)
		if data.Status==model.ActivitySaleOut{
			activityStatus.Status=model.ActivitySaleOut
		}else{
			activityStatus.Status=model.ActivityStart
		}
	//1.2没有开始
	}else if (now<data.StartTime){
		activityStatus.Status=model.ActivityNotStart
	}else{
		activityStatus.Status=model.ActivityEnd
	}

	//走到这里就有数据了,返回成功
	m=p.back(0,"success",data)
}

func (p *ProxyController)SecKill(){
	logs.Debug("enter sec info app")

	var m map[string]interface{} = make(map[string]interface{}, 16)
	defer func() {
		p.Data["json"] = m
		p.ServeJSON()
	}()

	product_id, err := p.GetInt("product_id", 0)
	if err != nil || product_id == 0 {
		m = p.back(1001, fmt.Sprintf("invalid product_id:%d", product_id), nil)
		logs.Error("invalid product id:%d", product_id)
		return
	}
	//从cookie中获得用户id
	userIdStr := p.Ctx.GetCookie("UserId")
	if len(userIdStr) == 0 {
		m = p.back(1003, fmt.Sprintf("invalid user_id:%s", userIdStr), nil)
		logs.Error("invalid user id:%d", product_id)
		return
	}
	//获取的用户id是字符串,转换成int
	//strings.parseInt()方法:string转int类型,可以设置进制和int64还是int32
	user_id, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil || user_id == 0 {
		m = p.back(1003, fmt.Sprintf("invalid user_id:%s", userIdStr), nil)
		logs.Error("invalid user id:%d", product_id)
		return
	}

	data, err := model.SecInfo(product_id, user_id)
	if err != nil {
		m = p.back(1002, fmt.Sprintf("service busy, product_id:%d, uid:%d",
			product_id, user_id), nil)
		logs.Error("service busy,product id:%d, err:%v user_id:%d", product_id, err, user_id)
		return
	}
	m = p.back(0, "success", data)
}