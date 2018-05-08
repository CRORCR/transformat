package balance

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterBalancer("random", &RandomBalance{})
}

type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	//1.如果是空切片,就直接返回
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	//2.根据长度,获得随机数,返回一个主机实例
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]

	return
}
