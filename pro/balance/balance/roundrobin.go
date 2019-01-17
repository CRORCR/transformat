package balance

import (
	"errors"
)

func init() {
	RegisterBalancer("roundrobin", &RoundRobinBalance{})
}

type RoundRobinBalance struct {
	curIndex int
}

func (p *RoundRobinBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	//1.如果是空切片,就直接返回
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}
	//2.如果角标越界,就从0重新开始
	lens := len(insts)
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	//3.角标加1
	//这里使用%和第二步判断长度是一样的,都是为了防止角标越界,都写也不冲突,删了也可以
	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens
	return
}
