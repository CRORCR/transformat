package model


import (
	"context"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var etcdClient *clientv3.Client
var productChan chan string

func initEtcd(conf *ModelConf) (err error) {

	etcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{conf.EtcdAddr},
		DialTimeout: 3 * time.Second,
	})

	if err != nil {
		logs.Warn("init etcd client failed, err:%v", err)
		return
	}
	//管道需要初始化
	productChan = make(chan string, 16)
	ctx, cancel := context.WithTimeout(context.Background(), 2 *time.Second)
	//   etcd_key /seckill/product/conf
	resp, err := etcdClient.Get(ctx, conf.EtcdProductKey)
	cancel()
	if err != nil {
		logs.Warn("get key %s failed, err:%v", conf.EtcdProductKey, err)
		return
	}
	//获得etcd最新配置,放入管道,给业务层用
	for _, ev := range resp.Kvs {
		logs.Debug(" %q : %q\n",  ev.Key, ev.Value)
		productChan <- string(ev.Value)
	}
	//以后就只要监听即可
	go WatchEtcd(conf.EtcdProductKey)
	return
}

func WatchEtcd(key string) {

	for {
		rch := etcdClient.Watch(context.Background(), key)
		for resp := range rch {
			for _, ev := range resp.Events {
				logs.Debug("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				//如果是put(有数据放入)就放入管道中
				if ev.Type == clientv3.EventTypePut {
					productChan <- string(ev.Kv.Value)
				}
			}
		}
	}
}
//提供获得管道方法
func GetProductChan() chan string {
	return productChan
}