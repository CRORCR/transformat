package config

type Notify interface{
	CallBack(*Config)
}
