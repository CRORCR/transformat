package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"

func main() {
	timestamp := GetTimestamp()               //1530784591
	timeStr := GetTimeStr()                   //2018-07-05 17:56:31
	fmt.Println(Timestamp2TimeStr(timestamp)) //2018-07-05 17:56:31
	fmt.Println(TimeStr2Timestamp(timeStr))   //1530784591
}

//获得时间戳
func GetTimestamp() int64 {
	return time.Now().Unix()
}

//获得当前时间(格式化)
func GetTimeStr() string {
	return time.Now().Format(TimeLayout)
}

//时间戳 转 string格式
func Timestamp2TimeStr(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(TimeLayout)
}

//string 转 时间戳
func TimeStr2Timestamp(timeStr string) int64 {
	// 得到本地时区 北京是东八区,如果不获得本地市区,得到的时间戳会领先8小时
	loc, _ := time.LoadLocation("Local")
	// 读取文件中的时间字符串末尾会带有换行符，需要将其去掉否则解析异常
	tm, error := time.ParseInLocation(TimeLayout, strings.Trim(timeStr, "\n"), loc)
	if error != nil {
		log.Fatal(error)
		return 0
	}
	return tm.Unix()
}
