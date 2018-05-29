package main

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	elastic "gopkg.in/olivere/elastic.v2"
)

var clients *elastic.Client

func initES(addr string) (err error) {
	clients, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.12.3:9200/"))
	if err != nil {
		logs.Error("connect to es error:%v", err)
		return
	}

	logs.Debug("connect es success")
	return
}

func Run(threadNum int) (err error) {

	go reload()

	for i := 0; i < threadNum; i++ {
		waitGroup.Add(1)
		go sendToEs()
	}
	waitGroup.Wait()
	return
}
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
	waitGroup.Done()
}
