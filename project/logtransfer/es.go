package main

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	elastic "gopkg.in/olivere/elastic.v2"
)

var clients *elastic.Client

func initES(addr string) (err error) {
	//访问etcd客户端   es只接受json格式   地址:"http://192.168.12.3:9200/"
	clients, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		logs.Error("connect to es error:%v", err)
		return
	}

	logs.Debug("connect es success")
	return
}

func Run(threadNum int) (err error) {

	go reload()
	//起一个线程发送就+1,发送完成-1
	for i := 0; i < threadNum; i++ {
		waitGroup.Add(1)
		go sendToEs()
	}
	//在这里等待全部发送
	waitGroup.Wait()
	return
}

//获得etcd中tpoic信息,去收集哪些日志
func reload() {
	for conf := range GetLogConfChan() {
		var topicArray []string
		err := json.Unmarshal([]byte(conf), &topicArray)
		if err != nil {
			logs.Error("unmarshal failed, err:%v conf:%s", err, conf)
			continue
		}
		reloadKafka(topicArray)
	}
}

//topic集合
func reloadKafka(topicArray []string) {
	for _, topic := range topicArray {
		kafka.AddTopic(topic)
	}
}

type EsMessage struct {
	Message string
}

func sendToEs() {
	for msg := range GetMessage() {
		var esMsg EsMessage
		esMsg.Message = msg.line
		logs.Debug("begin send to es succ")

		_, err := clients.Index().Index(msg.topic).Type(msg.topic).BodyJson(esMsg).Do()
		if err != nil {
			logs.Error("send to es failed, err:%v", err)
			continue
		}

		logs.Debug("send to es succ")
	}
	//发送完成,退出
	waitGroup.Done()
}
