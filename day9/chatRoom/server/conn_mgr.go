package server

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
	"transformat/day9/chatRoom/proto"
)
//聊天室需要管理所有请求链接,写一个管理的struct
type ClientMgr struct {
	//clientsMap维护所有客户端连接,关闭连接比较方法,根据key关闭
	clientsMap map[net.Conn]int
	//聊天室最大连接数
	maxClientNums int
	//msgChan用来保存客户端发送过来的消息
	msgChan chan *proto.Packet
	newClientChan chan net.Conn
	//链接关了,需要从管理中去掉
	closeChan chan net.Conn
	lock sync.RWMutex
}

func NewClientMgr(maxClients int) *ClientMgr {
	mgr :=  &ClientMgr {
		clientsMap:make(map[net.Conn]int, 1024),
		maxClientNums: maxClients,
		msgChan: make(chan *proto.Packet, 1000),
		newClientChan: make(chan net.Conn, 1000),
		closeChan: make(chan net.Conn, 1000),
	}

	go mgr.run()
	go mgr.procConn()
	return mgr
}

//如果新的连接进来,添加,并且有连接关闭,从map中去除
func (c *ClientMgr) procConn() {
	//记得加锁
	for {
		select {
		case conn := <- c.newClientChan:
			c.lock.Lock()
			c.clientsMap[conn] = 0
			c.lock.Unlock()
		case conn := <- c.closeChan:
			c.lock.Lock()
			delete(c.clientsMap, conn)
			c.lock.Unlock()
		}
	}
}

//遍历所有客户端发送过来的消息，并广播到所有的其他客户端
func (c *ClientMgr) run() {
	for msg := range c.msgChan {
		c.transfer(msg)
	}
}

//广播消息
func (c *ClientMgr)transfer(msg *proto.Packet) {
	//加读写锁
	c.lock.RLock()
	defer c.lock.RUnlock()
	for client, _ := range c.clientsMap {
		err := c.sendToClient(client, msg)
		if err != nil {
			continue
		}
	}
}

//发送消息给指定客户端
func (c *ClientMgr) sendToClient(client net.Conn, msg *proto.Packet) (err error) {
	return proto.WritePacket(client, msg.Cmd, msg.Body)
	/*
	断点续传,判断发送的数据是否全部发完
	var n int
	var sendBytes int
	msgLen := len(msg)
	for {
		n, err = client.Write(msg)
		if err != nil {
			fmt.Printf("send to client:%v failed, err:%v\n", client, err)
			client.Close()
			delete(c.clientsMap, client)
			return
		}
		sendBytes += n
		if sendBytes >= msgLen {
			break
		}
		msg = msg[sendBytes:]
	}
	return
	*/
}

//获得消息,放入管道
func (c *ClientMgr) addMsg(msg *proto.Packet) (err error) {
	//设置定时器,记得及时关闭
	ticker := time.NewTicker(time.Millisecond*10)
	defer ticker.Stop()
	//如果10毫秒,消息没有放入管道,就返回超时
	select {
	case c.msgChan <- msg:
		fmt.Printf("send to chan succ\n")
	case <- ticker.C:
		fmt.Printf("add msg timeout\n")
		err = errors.New("add msg timeout")
	}
	return
}