package balance

//传入一个切片 返回一个主机实例
//第二个参数是可变参数,随机和轮循用不上,这是给hash使用的
type Balancer interface {
	DoBalance([]*Instance, ...string) (*Instance, error)
}

