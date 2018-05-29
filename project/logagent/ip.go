package main

import "net"

var ipAarrays []string

func getIp() (err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, i := range interfaces {
		addrs, errRet := i.Addrs()
		if errRet != nil {
			err = errRet
			return
		}
		//获得所有的网关
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
				if ip.IsGlobalUnicast() {
					ipAarrays = append(ipAarrays, ip.String())
				}
			}
		}
	}
	return
}
