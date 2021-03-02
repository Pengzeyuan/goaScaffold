package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	//设置路由
	http.HandleFunc("/hello", Router)

	http.ListenAndServe("127.0.0.1:8089", nil)

}

func Router(resp http.ResponseWriter, request *http.Request) {
	resp.Write([]byte(GetLocalIp() + "/" + GetMac()))
}

//获取本机ip
func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get local ip failed")
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

//获取本机Mac
func GetMac() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("get local mac failed")
	}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr
		if mac.String() != "" {
			return mac.String()
		}
	}
	return ""
}
