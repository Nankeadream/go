package udp

import (
	"fmt"
	"log"
	"net"
)

func catchErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func UdpServer(buffer []byte) {
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	catchErr(err)
	fmt.Println(laddr)
	// 1 sockert 2 bind
	server, err := net.ListenUDP("udp", laddr)
	catchErr(err)
	defer server.Close()
	//不需要accept
	//3 收发
	n, raddr, err := server.ReadFromUDP(buffer)
	fmt.Println(err)
	msg := fmt.Sprintf("Client=%v,data=%s", raddr, string(buffer[:n]))
	fmt.Println(msg)
	server.WriteToUDP([]byte(msg), raddr)
}
