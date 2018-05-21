package main

//轮循
type RoundBalance struct {
	//记录下标
	curIndex int
}

func (r *RoundBalance) DoBalance(addrList []string) string {
	//对长度取余,确保在自增的时候,维持在长度范围内
	l := len(addrList)
	r.curIndex = r.curIndex % l

	addr := addrList[r.curIndex]
	r.curIndex++
	return addr
}
