package main

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"transformat/day7/demo1_balance/balance"
)

//实现一致性hash负载均衡算法
//一致性hash 每次过来都是访问同一台ip
type HashBalance struct {
}

func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	//如果有实例过来,就赋值上,如果没有就使用随机的
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("No backend instance")
		return
	}
	//获得一个hash值
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey), crcTable)
	//取余
	index := int(hashVal) % lens
	inst = insts[index]
	return
}
