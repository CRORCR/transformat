package main

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	filename:="./my.log"
	tails,err:=tail.TailFile(filename,tail.Config{
		ReOpen:true,
		Follow:true,
		Location:&tail.SeekInfo{Offset:0,Whence:2},
		MustExist:false,
		Poll:true,
	})

	if err!=nil{
		fmt.Println("tail is failed ,err ",err)
		return
	}

	var msg *tail.Line
	var ok bool
	for true {
		msg,ok=<-tails.Lines
		if !ok{
			fmt.Printf("tail file close reopen ,filename:%s",tails.Filename)
			time.Sleep(100*time.Millisecond)
			continue
		}
		fmt.Printf("msg:",msg.Text)
	}

}

