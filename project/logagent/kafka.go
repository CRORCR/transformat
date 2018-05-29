package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

//发送给kafka的消息,必须按照我的定义来
type Message struct {
	line  string
	topic string
}

type KafkaSender struct {
	client   sarama.SyncProducer
	lineChan chan *Message
}

var kafkaSender *KafkaSender

func initKafka() {
	kafkaSender = NewKafka()
	return
}
func NewKafka() (kafka *KafkaSender) {
	kafka = &KafkaSender{lineChan: make(chan *Message, 10000)}
	//新建一个配置
	config := sarama.NewConfig()
	//设置是否需要回报
	config.Producer.RequiredAcks = sarama.NoResponse
	//设置分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//需要返回不
	config.Producer.Return.Successes = true

	//配置文件弄好,就可以生成一个client,然后放入kafka中,从配置文件读取线程数量,去发送消息到kafka
	kafkaAdr := appConfig.kafkaAddr
	asyncclient, err := sarama.NewSyncProducer([]string{kafkaAdr}, config)
	if err != nil {
		panic("init kafka failed")
		logs.Warn("init kafka client failed,err:%v", err)
		return
	}
	kafka.client = asyncclient
	for i := 0; i < appConfig.KafkaThreadNum; i++ {
		go kafka.sendToKafka()
	}
	return
}

//起多个线程去往kafka发送数据
func (k *KafkaSender) sendToKafka() {
	for v := range k.lineChan {
		//格式化消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = v.topic
		msg.Value = sarama.StringEncoder(v.line)
		_, _, err := k.client.SendMessage(msg)
		if err != nil {
			logs.Error("send message to kafka failed,err:%v\n", err)
			return
		}
	}
}

func (k *KafkaSender) addMessage(topic, line string) (err error) {
	k.lineChan <- &Message{line: line, topic: topic}
	return
}
