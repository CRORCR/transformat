package proto

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const (
	CmdLoginRequest       = 1001
	CmdLoginReponse       = 1002
	CmdReginsterRequest   = 1003
	CmdReginsterReponse   = 1004
	CmdSendMessageRequest = 1005
	CmdSendMessageReponse = 1006
	CmdBroadMessage       = 1007
)

type ResponseBase struct {
	Error   int    `json:"errno"`
	Message string `json:"message"`
}
type LoginResponse struct {
	ResponseBase
}

type MessageReposne struct {
	ResponseBase
}

type BroadMessage struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

type MessageRequest struct {
	Message  string `json:"message"`
	UserName string `json:"username"`
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `josn:"password"`
}
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Sex      string `json:"sex"`
}

type ReginsterReposne struct {
	ResponseBase
}
type Packet struct {
	Cmd  int32
	Body []byte
}

func ReadPacket(conn net.Conn) (body []byte, cmd int32, err error) {
	var length int32
	err = binary.Read(conn, binary.BigEndian, &length)
	if err != nil {
		fmt.Printf("read from conn:%v failed ,err:%v \n", conn, err)
		return
	}
	fmt.Printf("read cmd seccess:%d \n", cmd)
	var buf []byte = make([]byte, length)
	_, err = io.ReadFull(conn, buf)
	if err != nil {
		fmt.Printf("read body from conn %v failed,err:%v \n", conn, err)
		return
	}
	body = buf
	fmt.Printf("read body success \n", string(buf))
	return
}

func WritePacket(conn net.Conn, cmdno int32, body []byte) (err error) {
	var length int32 = int32(len(body))
	err = binary.Write(conn, binary.BigEndian, length)
	if err != nil {
		fmt.Printf("write length is failed err:%v \n", err)
		return
	}
	fmt.Printf("write length is success \n")

	err = binary.Write(conn, binary.BigEndian, cmdno)
	if err != nil {
		fmt.Printf("write cmdno failed err:%v \n", err)
		return
	}
	var n int
	var sendBytes int
	msgLen := len(body)
	for {
		n, err := conn.Write(body)
		if err != nil {
			fmt.Printf("send to client %v failed,err:%v", conn, err)
			return
		}
		sendBytes += n
		if sendBytes >= msgLen {
			break
		}
		body = body[sendBytes:]
	}
	return
}
