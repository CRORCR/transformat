package main

import (
	"net"
)

var ipArrays []string

//在初始化etcd之前拿到本地ip
//根据本地网卡接口 如果是ipv4地址就加到切片里面
func getLocalIP() (err error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, i := range ifaces {
		addrs, errRet := i.Addrs()
		if errRet != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				if ip.IsGlobalUnicast() {
					ipArrays = append(ipArrays, ip.String())
				}
			}
		}
	}
	return
}
