package main

import (
	"fmt"
	"net"
)

func runServer(l net.Listener) (err error) {
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("accept is failed ,err:%v \n", err)
			continue
		}
		go process(conn)
	}

}

func process(conn net.Conn) {

}
