package client

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"transformat/day9/proto"
)

var recvMeg chan interface{}

func main() {
	conn, err := net.Dial("tcp", "192.168.14.200:18080")
	if err != nil {
		fmt.Printf("dial server failed ,err %v\n", err)
		return
	}
	recvMeg = make(chan interface{}, 1000)
	defer conn.Close()
	go read(conn)
	msg := <-recvMeg
	loginResp, ok := msg.(*proto.LoginResponse)
	if !ok {
		fmt.Printf("unexpect msg:%T,%+v \n", msg, msg)
		return
	}
	if loginResp.Error != 0 {
		fmt.Printf("login failed ,err:%v \n", err)
		return
	}
	fmt.Println("login success")
	for {
		var data string
		reader := bufio.NewReader(os.Stdin)
		data, err := reader.ReadString('\n')
		if err != nil {
			continue
		}
		err = sendMessage(conn, data)
		if err != nil {
			fmt.Printf("send message failed ,err:%v \n", err)
			return
		}
	}
}

func read(conn net.Conn) {
	for {
		body, cmd, err := proto.ReadPacket(conn)
		if err != nil {
			fmt.Printf("read from server failed,err:%v\n", err)
			return
		}
		switch cmd {
		case proto.CmdLoginReponse:
			err = processLoginReponse(conn, body)
		case proto.CmdSendMessageReponse:
			err = processSendMessageResponse(conn, body)
		case proto.CmdBroadMessage:
			err = processBroadMessage(conn, body)
		default:
			fmt.Printf("unsupport cmd[%v] \n", cmd)
			return
		}
	}
}

func processLoginReponse(conn net.Conn, body []byte) (err error) {
	var loginReponse proto.LoginResponse
	err = json.Unmarshal(body, &loginReponse)
	if err != nil {
		fmt.Printf("unmarshal failed,err:%v \n", err)
		return
	}
	recvMeg <- &loginReponse
	return
}

func processSendMessageResponse(conn net.Conn, body []byte) (err error) {
	var sendMessage proto.MessageReposne
	err = json.Unmarshal(body, &sendMessage)
	if err != nil {
		fmt.Printf("unmarshal is failed ,err:%v \n", err)
		return
	}
	if sendMessage.Error != 0 {
		fmt.Printf("消息发送失败:%v \n", sendMessage.Message)
		return
	}
	return
}

func processBroadMessage(conn net.Conn, body []byte) (err error) {
	var msg proto.BroadMessage
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Printf("unmashal is failed err:%v \n", err)
		return
	}
	fmt.Printf("%s:\n %s:\n\n", msg.Message, msg.Username)
	return
}

func sendMessage(conn net.Conn, data string) (err error) {
	var message proto.MessageRequest
	message.Message = data
	message.UserName, _ = os.Hostname()

	body, err := json.Marshal(message)
	if err != nil {
		fmt.Printf("marshal is failed ,err:%v \n", err)
		return
	}
	err = proto.WritePacket(conn, proto.CmdSendMessageRequest, body)
	if err != nil {
		fmt.Printf("send to server failed ,err;5v \n", err)
		return
	}
	return
}

func login(conn net.Conn) (err error) {
	var loginReq proto.LoginRequest
	loginReq.UserName = "admin"
	loginReq.Password = "admin"

	body, err := json.Marshal(loginReq)
	if err != nil {
		fmt.Printf("marshal is failed err:%v \n", err)
		return
	}
	err = proto.WritePacket(conn, proto.CmdSendMessageRequest, body)
	if err != nil {
		fmt.Printf("send to server failed ,err;5v \n", err)
		return
	}
	return
}
