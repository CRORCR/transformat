package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Config struct{
	fileName string
	lastModifyTime  int64
	data map[string]string
	rwLock sync.RWMutex
	notifyList []Notify
}

func NewConfigG(filename string)(config *Config,err error){
	config=&Config{fileName:filename,data:make(map[string]string,1024),}
	m,err:=config.parse()
	if err!=nil{
		fmt.Println("转换失败")
		return
	}
	config.rwLock.Lock()
	config.data=m
	config.rwLock.Unlock()

	go config.reload()
	return
}

func(config *Config)parse()(m map[string]string,err error){
	m=make(map[string]string,1024)
	//打开文件
	file,err:=os.Open(config.fileName)
	if err!=nil{
		fmt.Println("打开文件有误")
		return
	}
	//输出错误的序列号
	var index int
	reader:=bufio.NewReader(file)
	for{
		line,errRet:=reader.ReadString('\n')
		if errRet==io.EOF{
			break //读取到末尾
		}
		if errRet!=nil{
			err=errRet//如果不是读取到末尾的错误,就返回
			return
		}
		index++
		//去除空格
		line=strings.TrimSpace(line)
		//如果是空格 或者以...开始都循环下一次
		if len(line)==0 || line[0]=='\n' || line[0]=='#' || line[0]==';'{
			continue
		}
		arr:=strings.Split(line,"=")
		if len(arr)==0{
			fmt.Printf("invalid config,line:%d \n",index)
			continue
		}
		key:=strings.TrimSpace(arr[0])
		if len(key)==0{
			fmt.Printf("invalid config,line:%d \n",index)
			return
		}
		if len(arr)!=2{
			continue
		}
		value:=strings.TrimSpace(arr[1])
		m[key]=value
	}
	return
}

func (config *Config)reload(){
	ticker:=time.NewTicker(time.Second*5)
	for _=range ticker.C{
		func(){
			file,_:=os.Open(config.fileName)
			defer file.Close()

			fileInfo,_:=file.Stat()
			curModifyTime:=fileInfo.ModTime().Unix()
			if curModifyTime>config.lastModifyTime{
				m,_:=config.parse()
				config.rwLock.Lock()
				config.data=m
				config.rwLock.Unlock()

				config.lastModifyTime=curModifyTime
				for _,n :=range config.notifyList{
					n.CallBack(config)
				}
			}
		}()
	}
}
func (config *Config)AddNotifyer(n Notify){
	config.notifyList=append(config.notifyList,n)
}

func (config *Config)GetInt(key string)(value int,err error){
	config.rwLock.Lock()
	defer config.rwLock.Unlock()

	str,ok:=config.data[key]
	if !ok{
		err=fmt.Errorf("key[%s] not found",key)
		return
	}
	value,_=strconv.Atoi(str)
	return
}

func(config *Config)GetIntDefault(key string,def int)(value int){
	config.rwLock.Lock()
	defer config.rwLock.Unlock()

	str,ok:=config.data[key]
	if !ok{
		value=def
		return
	}

	value,err:=strconv.Atoi(str)
	if err!=nil{
		value=def
		return
	}
	return
}

func (c *Config) GetString(key string)(value string, err error) {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	value, ok := c.data[key]
	if !ok {
		err = fmt.Errorf("key[%s] not found", key)
		return
	}

	return
}