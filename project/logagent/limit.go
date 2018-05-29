package main

import (
	"github.com/astaxie/beego/logs"
	"sync/atomic"
	"time"
)

//每秒限流
//日志收集其实是辅助性程序,业务占用cpu比较大的时候,日志收集应该限流
//一秒收集多少日志 进行限流
type SecondLimit struct {
	unixSecond int64 //当前秒数
	curCount   int32 //收集了多少日志
	limit      int32 //日志上限
}

func NewSecondLimit(limit int32) *SecondLimit {
	secLimit := &SecondLimit{
		unixSecond: time.Now().Unix(),
		curCount:   0,
		limit:      limit,
	}

	return secLimit
}

func (s *SecondLimit) Add(count int) {
	sec := time.Now().Unix()
	//等于 当前秒数   统计日志(原子操作,相加)
	if sec == s.unixSecond {
		atomic.AddInt32(&s.curCount, int32(count))
		return
	}
	//更新计数 清零
	atomic.StoreInt64(&s.unixSecond, sec)
	atomic.StoreInt32(&s.curCount, int32(count))
}

func (s *SecondLimit) Wait() bool {
	for {
		sec := time.Now().Unix()
		//同一秒,超过上限,等待一毫秒 循环
		if sec == atomic.LoadInt64(&s.unixSecond) && s.curCount >= s.limit {
			time.Sleep(time.Millisecond)
			logs.Debug("limit is running, limit:%d s.curCount:%d ", s.limit, s.curCount)
			continue
		}
		//重新计数
		if sec != atomic.LoadInt64(&s.unixSecond) {
			atomic.StoreInt64(&s.unixSecond, sec)
			atomic.StoreInt32(&s.curCount, 0)
		}
		logs.Debug("limit is exited")
		return false
	}
}
