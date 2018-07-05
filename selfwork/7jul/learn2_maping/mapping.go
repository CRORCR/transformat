package main

import (
	"fmt"
	"sort"
)

var intMap = make(map[int]int)

//全局map会把内存被分配到堆上面
func main() {
	demo1()
	demo2()
}

func demo1() {
	intMap[1] = 11
	intMap[2] = 111
	//1.它修改了当前key的标记，而不是直接删除了内存数据
	delete(intMap, 1)

	//2.map删除了,但是还是会占用内存
	for k, _ := range intMap {
		delete(intMap, k)
	}
	//3.彻底清空nil,然后等待垃圾回收处理
	intMap = nil

	//4.为什么delete删除了元素,但是没有真正删除元素?
	//这么设计是为了在删除元素后,再遍历不会发送错乱的情况,删除的同时读取会发送异常
	//而且删除了这个元素,下次k相同的元素再次进入,会覆盖之前的元素
}

//二维数组演示
func demo2() {
	logMap := make(map[string]map[string]string)
	mm := make(map[string]string)
	mm["gotreply"] = "4"
	mm["gotreply"] = "2"
	mm2 := make(map[string]string)
	mm2["findnode"] = "xxxx"
	mm2["pingpong"] = "pingpong"
	logMap["192.168.3.100"] = mm
	logMap["192.168.3.100"] = mm2
	fmt.Println(logMap) //map[192.168.3.100:map[findnode:xxxx pingpong:pingpong]]
}

//map排序(k) 先取出所有K,然后对k排序,再变了k输出value结果,达到排序效果
func demo3() {
	var m = map[string]string{
		"3GotReply2":     "ptype:2 from:4f554827437825cc nowTime:1530584974274",
		"5GotReply4":     "ptype:4 from:4f554827437825cc nowTime:1530584977777",
		"2WritePingTime": "TO:4f554827437825cc NowTime:1530584974270\n",
		"4findNode":      "TO:4f554827437825cc NowTime:1530584977770",
		"1ip":            "192.168.3.123:30303",
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
