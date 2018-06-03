package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	//新建一个配置
	config := sarama.NewConfig()
	//发送消息是否等待ack,是否处理好
	config.Producer.RequiredAcks = sarama.WaitForAll
	//分区,如果存一个分区,并发性能不好,kafka是分布式的,最后是多个分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//创建一条消息
	msg := &sarama.ProducerMessage{}
	//topic主要区分不同的业务
	msg.Topic = "nginx_log"
	//实在的消息
	msg.Value = sarama.StringEncoder("this is a good test, my message is good")

	//创建同步的客户端
	client, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()
	//发送消息 返回三个信息 消息所在分区/偏移量
	for {
		pid, offset, err := client.SendMessage(msg)

		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		//5毫秒发送一次,后台可以一直读取
		time.Sleep(5 * time.Millisecond)
	}
}
