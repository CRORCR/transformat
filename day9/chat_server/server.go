package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"transformat/day9/proto"
)

func runServer(l net.Listener) (err error) {
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("connection is failed err:%v \n", err)
			continue
		}
		clientMgr.newClientChan <- conn
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer func() {
		clientMgr.closeChan <- conn
		conn.Close()
	}()
	for {
		body, cmd, err := proto.ReadPacket(conn)
		if err != nil {
			fmt.Printf("read from conn failed err:%v \n", err)
			return
		}
		err = processRequest(conn, body, cmd)
		if err != nil {
			fmt.Printf("processRequest[%v] failed ,err:%v \n", cmd, err)
			return
		}
	}
}

func processRequest(conn net.Conn, body []byte, cmd int) (err error) {
	switch cmd {
	case proto.CmdLoginRequest:
		err = processLogin(conn, body)
	case proto.CmdReginsterRequest:
		err = processRegister(conn, body)
	case proto.CmdSendMessageRequest:
		err = processMessage(conn, body)
	default:
		fmt.Printf("unsupport cmd[%v] \n", cmd)
		err = errors.New("unsupport cmd")
		return
	}
	return
}

func processMessage(conn net.Conn, body []byte) (err error) {
	var mess proto.MessageRequest
	err = json.Unmarshal(body, &mess)
	if err != nil {
		fmt.Printf("unmrshal failed[%v] \n", err)
		return
	}
	var broad proto.BroadMessage
	broad.Message = mess.Message
	broad.Username = mess.UserName

	data, err := json.Marshal(broad)
	if err != nil {
		fmt.Printf("marshal is failed err:%v \n", err)
		return
	}
	packet := &proto.Packet{
		Cmd:  proto.CmdBroadMessage,
		Body: data,
	}
	clientMgr.addMsg(packet)
	return

	return
}

func processLogin(conn net.Conn, body []byte) (err error) {
	var log proto.LoginRequest
	err = json.Unmarshal(body, &log)
	if err != nil {
		fmt.Printf("unmarshal is failed err:%v \n", err)
		return
	}
	var logresp proto.LoginResponse
	logresp.Error = 100
	logresp.Message = "username or password not right"

	if log.UserName == "admin" && log.Password == "admin" {
		logresp.Error = 0
		logresp.Message = "success"
	}
	data, err := json.Marshal(logresp)
	if err != nil {
		fmt.Printf("mrshal is failed ,err:%v \n", err)
		return
	}
	return proto.WritePacket(conn, proto.CmdLoginReponse, data)
}

func processRegister(conn net.Conn, body []byte) (err error) {
	return
}
