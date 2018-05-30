package config

//需要热加载的应用只要实现这个接口就可以了,然后加到需要通知的切片里面(notifyList)
//文件有变更就会通知对应的应用

type Notifyer interface {
	Callback(*Config)
}
