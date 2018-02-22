package main

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
	"transformat/day9/proto"
)

type ClientMgr struct {
	//clientsMap维护所有客户端连接
	clientsMap    map[net.Conn]int
	maxClientNums int
	//msgChan用来保存客户端发送过来的消息
	msgChan       chan *proto.Packet
	newClientChan chan net.Conn
	closeChan     chan net.Conn
	lock          sync.RWMutex
}

func NewClientMgr(maxClients int) *ClientMgr {
	mgr := &ClientMgr{
		clientsMap:    make(map[net.Conn]int, 1024),
		maxClientNums: maxClients,
		msgChan:       make(chan *proto.Packet, 1000),
		newClientChan: make(chan net.Conn, 1000),
		closeChan:     make(chan net.Conn, 1000),
	}
	go mgr.run()
	go mgr.procConn()
	return mgr
}

//循环所有客户端发过来的消息,并广播到所有其他客户端
func (c *ClientMgr) run() {
	for msg := range c.msgChan {
		c.transfer(msg)
	}
}

//广播消息
func (c *ClientMgr) transfer(msg *proto.Packet) {
	c.lock.Lock()
	defer c.lock.Unlock()

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
}

func (c *ClientMgr) procConn() {
	for {
		select {
		case conn := <-c.newClientChan:
			c.lock.Lock()
			defer c.lock.Unlock()
			c.clientsMap[conn] = 0
		case conn := <-c.closeChan:
			c.lock.Lock()
			defer c.lock.Unlock()
			delete(c.clientsMap, conn)
		}
	}
}

func (c *ClientMgr) addMsg(msg *proto.Packet) (err error) {
	ticker := time.NewTicker(time.Microsecond * 10)
	defer ticker.Stop()

	select {
	case c.msgChan <- msg:
		fmt.Printf("send to chan success \n")
	case <-ticker.C:
		fmt.Printf("add msg timeOut...")
		err = errors.New("add msg timeOut...")
	}
	return

}
