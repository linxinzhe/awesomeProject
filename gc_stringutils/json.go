package gc_stringutils

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"serverName"`  //通过struct tag定义来实现小写的首字母
	ServerIP   string `json:"serverIP"`
}

type ServerSlice struct {
	Servers []Server
}

func Unmarshal() {
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	var serverSlice ServerSlice
	json.Unmarshal([]byte(str), &serverSlice) // 参数interface{}等于任意类型 注意是serverSlice的地址

	fmt.Println(serverSlice)
}

func Marshall() {
	var s ServerSlice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
