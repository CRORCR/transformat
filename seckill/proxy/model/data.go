package model

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
)

const(
	ActivityStart    = 1
	ActivityEnd      = 2
	ActivitySaleOut  = 3
	ActivityNotStart = 4
)
//秒杀结果结构体
type SecKillResult struct {
	UserId    int
	ProductId int
	//资格(抢购成功)
	Token  string
	Status int
}
//秒杀抢购请求
type SecKillRequest struct {
	UserId     int
	ProductId  int
	UserIp     string
	resultChan chan *SecKillResult
}

//商品使用map保存,秒杀的商品不确定,随机的
//这里的map是全局的,所以需要加锁(防止正在抢购中修改配置)
//这里的map读比较多,写很少(读写锁)
type SecProxyData struct {
	activityMap map[int]*Activity
	rwLock sync.RWMutex
	requestChan chan *SecKillRequest
}

//定义活动结果 商品id 开始结束时间 状态码
type Activity struct{
	ProductId int    `json:"product_id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64	 `json:"end_time"`
	Status    int	 `json:"status"`
}

//定义商品全局变量,从管道获取之后,赋值
var secProxyData *SecProxyData

//把活动放入map中的方法
//在这里加锁
func(s *SecProxyData)UpdateActive(prodyuctArry []*Activity)(err error){
	s.rwLock.Lock()
	defer s.rwLock.Unlock()

	//1.循环所有活动,加入map中
	for _,v :=range prodyuctArry{
		s.activityMap[v.ProductId]=v
	}
	//2.所有的活动必须在map中都存在
	//如果最新配置中没有活动,map中也要删除
	for k,_:=range s.activityMap{
		found:=false
		for _,v:=range prodyuctArry{
			if k==v.ProductId{
				found=true
			}
		}
		if found==false{
			delete(s.activityMap,k)
		}
	}
	return
}

func(s *SecProxyData)Reload(){
	//1.遍历管道,如果有数据,就更新
	for productConf:=range GetProductChan(){
		var activityArr []*Activity
		err:=json.Unmarshal([]byte(productConf),&activityArr)
		if err!=nil{
			logs.Error("unmarshal failed,err:%v,conf:%s",err,productConf)
			continue
		}
		//2.有数据,调用update
		err = s.UpdateActive(activityArr)
		if err != nil {
			logs.Error("UpdateActivity failed, err:%v, conf:%s", err, productConf)
			continue
		}

		logs.Debug("reload conf from etcd succ, new conf:%s", productConf)
	}
	return
}

func (s *SecProxyData)GetActive(productId int)(a *Activity,err error){
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	a,ok:=s.activityMap[productId]
	if !ok{
		err=fmt.Errorf("product:%s not found activity,err:%v",productId,err)
		return
	}
	return
}

//chan是线程安全的,不需要加锁
//需要设置超时时间,否则会出现阻塞
func(s *SecProxyData)AddRequest(req *SecKillRequest)(err error){
	timer:=time.NewTicker(2*time.Second)
	defer func(){
		timer.Stop()
	}()
	select{
	case s.requestChan<-req:
	case <-timer.C:
		err=fmt.Errorf("time out")
		return
	}
	return
}
