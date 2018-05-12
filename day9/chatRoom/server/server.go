package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"transformat/day9/chatRoom/proto"
)

func runServer(l net.Listener) (err error) {
	fmt.Println("run server succ")
	for {
		conn, _ := l.Accept()
		//有消息,就放入管理消息中
		clientMgr.newClientChan <- conn
		//获得请求,启动goroute处理
		go process(conn)
	}
}

func process(conn net.Conn) {
	//关闭连接,并从管理中心map删除当前连接
	defer func() {
		clientMgr.closeChan <- conn
		conn.Close()
	}()

	for {
		body, cmd, err :=  proto.ReadPacket(conn)
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			return
		}
		//处理请求
		err = processRequest(conn, body, cmd)
		if err != nil {
			fmt.Printf("processRequest[%v] failed, err:%v\n", cmd, err)
			return
		}
		/*
		var buf []byte = make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			return
		}

		buf = buf[0:n]
		clientMgr.addMsg(buf)
		*/
	}
}

func processRequest(conn net.Conn, body []byte, cmd int32) (err error) {
	//判断是什么请求
	switch cmd {
	case proto.CmdLoginRequest:
		err = processLogin(conn, body)
	case proto.CmdRegisterRequest:
		err = processRegister(conn, body)
	case proto.CmdSendMessageRequest:
		err = processMessage(conn, body)
	default:
		//不支持的请求
		fmt.Printf("unsupport cmd[%v] \n", cmd)
		err = errors.New("unsupport cmd")
		return
	}
	return
}

func processLogin(conn net.Conn, body []byte) (err error) {

	fmt.Printf("begin process login request\n")
	var loginRequest proto.LoginRequest
	//1.反序列化
	err = json.Unmarshal(body, &loginRequest)
	if err != nil {
		fmt.Printf("Unmarshal failed[%v]\n", err)
		return
	}
	fmt.Printf(" process login request：%+v\n", loginRequest)
	var loginResp proto.LoginResponse
	loginResp.Errno = 100
	loginResp.Message = "username or password not right"
	//2.校验逻辑
	if loginRequest.Username == "admin" &&loginRequest.Password == "admin" {
		loginResp.Errno = 0
		loginResp.Message = "success"
	}
	//3.返回回报
	data, err := json.Marshal(loginResp)
	if err != nil {
		fmt.Printf("Marshal failed[%v]\n", err)
		return
	}
	fmt.Printf(" write login response %+v\n", loginResp)
	//3.写自定义流数据
	return proto.WritePacket(conn, proto.CmdLoginResponse, data)
}

func processRegister(conn net.Conn, body []byte) (err error) {
	//自己实现去
	return
}

func processMessage(conn net.Conn, body []byte) (err error) {
	fmt.Printf("begin process login request\n")
	var messageReq proto.MessageRequest
	err = json.Unmarshal(body, &messageReq)
	if err != nil {
		fmt.Printf("Unmarshal failed[%v]\n", err)
		return
	}
	var broadMessage proto.BroadMessage
	broadMessage.Message = messageReq.Message
	broadMessage.Username = messageReq.Username
	body, err = json.Marshal(broadMessage)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}
	packet := &proto.Packet {
		Cmd: proto.CmdBroadMessage,
		Body: body,
	}
	//去广播消息
	clientMgr.addMsg(packet)
	return
}
