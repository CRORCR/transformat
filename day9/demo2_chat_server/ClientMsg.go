package main

import (
	"net"
	"sync"
)

type ClientMsg struct {
	ClientMap  map[net.Conn]int
	MaxNum     int
	MsgChan    chan []byte
	clientChan chan net.Conn
	lock       sync.RWMutex
}
