package proto

import (
	"net"
	"encoding/binary"
	"fmt"
	"io"
)

//cmdno常量
const (
	CmdLoginRequest = 1001
	CmdLoginResponse = 1002
	CmdRegisterRequest = 1003
	CmdRegisterResponse = 1004
	CmdSendMessageRequest = 1005
	CmdSendMessageResponse = 1006
	CmdBroadMessage = 1007 //广播
)
//登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	ResponseBase
}
type ResponseBase struct {
	Errno int `json:"errno"`
	Message string `json:"message"`
}
//注册请求
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Sex string `json:"sex"`
}

type RegisterResponse struct {
	ResponseBase
}
//发送消息 消息和用户(谁发的)
type MessageRequest struct {
	Message string `json:"message"`
	Username string `json:"username"`
}

type BroadMessage struct {
	Message string `json:"message"`
	Username string `json:"username"`
}

type MessageResponse struct {
	ResponseBase
}

type Packet struct {
	Cmd int32
	Body []byte
}

//读写标准是自己定义的:长度 协议号 消息(读写都是这个格式)
//前四个字节：length
//4个字节：   cmdno  协议号(登录/注册/发消息)
//body:      []byte
func ReadPacket(conn net.Conn) (body []byte, cmd int32, err error) {
	var length int32
	err = binary.Read(conn, binary.BigEndian, &length)
	if err != nil {
		fmt.Printf("read from conn:%v failed, err:%v\n", conn, err)
		return
	}
	fmt.Printf("read length succ:%d\n", length)

	err = binary.Read(conn, binary.BigEndian, &cmd)
	if err != nil {
		fmt.Printf("read from conn:%v failed, err:%v\n", conn, err)
		return
	}
	fmt.Printf("read cmd succ:%d\n", cmd)

	var buf []byte = make([]byte, length)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		fmt.Printf("read body from conn %v failed, err:%v\n", conn, err)
		return
	}
	body = buf
	fmt.Printf("read body succ:%v\n", string(buf))
	return
	/*
	断点读取 == _, err = io.ReadFull(conn, buf) 效果一样
	var curReadBytes int32 = 0
	for {
		n, errRet := conn.Read(buf)
		if errRet != nil {
			err = errRet
			fmt.Printf("read body from conn %v failed, err:%v\n", conn, err)
			return
		}

		body = append(body, buf[0:n]...)
		curReadBytes += int32(n)
		if (curReadBytes == length) {
			break
		}

		buf = make([]byte, length - curReadBytes)
	}
	return*/
}

//前四个字节：length
//4个字节：   cmdno
//body:      []byte

//BigEndian 网络传输有大端和小端存储(正序和反序),一般都是使用大端
//aa    cc
//bb    bb
//cc    aa
func WritePacket(conn net.Conn, cmdno int32, body []byte) (err error) {
	var length int32 = int32(len(body))
	//binary二进制读写操作
	//1.写入长度
	err = binary.Write(conn, binary.BigEndian, length)
	if err != nil {
		fmt.Printf("write length failed, err:%v\n", err)
		return
	}
	fmt.Printf("write length succ:%d\n", length)
	//写入协议号
	err = binary.Write(conn, binary.BigEndian, cmdno)
	if err != nil {
		fmt.Printf("write cmd no failed, err:%v\n", err)
		return
	}
	fmt.Printf("write cmdno succ:%d\n", cmdno)
	//写入body
	//断点续传,判断发送的数据是否全部发完
	var n int
	var sendBytes int
	msgLen := len(body) //消息长度
	for {
		n, err = conn.Write(body)
		if err != nil {
			fmt.Printf("send to client:%v failed, err:%v\n", conn, err)
			return
		}
		sendBytes += n
		//判断是否全部发送,不是-->续传
		if sendBytes >= msgLen {
			break
		}
		//续传
		body = body[sendBytes:]
	}
	fmt.Printf("write body succ:%v\n", string(body))
	return
}