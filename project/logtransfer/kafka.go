package main

import (
	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

type LogData struct {
	line  string
	topic string
}
type KafkaObj struct {
	consumer sarama.Consumer
	topic    string
}

type KafkaMgr struct {
	topicMap  map[string]*KafkaObj
	kafkaAddr string
	msgChan   chan *LogData
}

var kafka *KafkaMgr

func initKafka() {
	kafka = NewKafKaMgr()
}
func NewKafKaMgr() (kafka *KafkaMgr) {
	kafkaAddr := appConfig.kafkaAddr
	kafka = &KafkaMgr{
		topicMap:  make(map[string]*KafkaObj, 10000),
		kafkaAddr: kafkaAddr,
		msgChan:   make(chan *LogData, 11),
	}
	return
}

func (k *KafkaMgr) AddTopic(topic string) {
	obj, ok := k.topicMap[topic]
	if ok {
		return
	}
	//如果不存在,就添加
	obj = &KafkaObj{
		topic: topic,
	}
	//消费端
	consumer, err := sarama.NewConsumer([]string{k.kafkaAddr}, nil)
	if err != nil {
		logs.Error("failed to connection kafka")
		return
	}
	obj.consumer = consumer
	//指定消费那个topic kafka是分布式的,可能存在多个机器上,根据每个分片启动线程去读取
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		logs.Error("failed to get list of partitions,err:%v", err)
		return
	}
	for partition := range partitions {
		pc, errRet := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if errRet != nil {
			err = errRet
			return
		}
		//一个分片起一个线程去消费
		go func(p sarama.PartitionConsumer) {
			for msg := range p.Messages() {
				logs.Debug("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				logData := &LogData{
					line:  string(msg.Value),
					topic: msg.Topic,
				}
				k.msgChan <- logData
			}
		}(pc)
	}
}

func GetMessage() chan *LogData {
	return kafka.msgChan
}
