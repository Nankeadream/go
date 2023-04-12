package main

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"time"
)

// ClientRead 3 收发数据
func ClientRead(conn *net.TCPConn, buffer []byte, exit chan<- struct{}) {
	for {
		err := conn.SetReadDeadline(time.Now().Add(time.Second)) //每次等待一秒就超时
		catchErr(err)
		n, err := conn.Read(buffer)
		if err != nil {
			if _, ok := err.(*net.OpError); !ok {
				exit <- struct{}{}
				return
			}
			continue
		}
		fmt.Printf("从服务器%v接收到数据%v\n", conn.RemoteAddr(), string(buffer[:n]))
	}

}
func ClientSend(conn *net.TCPConn, exit chan<- struct{}) {
	var input string
	for {
		fmt.Println("请输入要发送的数据:")
		fmt.Scanln(&input)
		fmt.Printf("发送数据%v\n", input)
		if input == "exit" {
			exit <- struct{}{}
			return
		}
		_, err := conn.Write([]byte(input))
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
func catchErr(err error) {
	if err != nil {
		log.Fatalln(err)
		return
	}
}
func main() {
	raddr, err := net.ResolveTCPAddr("tcp", "192.168.3.11:9999")
	catchErr(err)
	//1 创建socket 2 发起对服务器端的连接
	conn, err := net.DialTCP("tcp", nil, raddr)
	catchErr(err)
	fmt.Println("客户端启动", conn.LocalAddr())
	//4 关闭
	defer conn.Close()
	exit := make(chan struct{})
	buffer := make([]byte, 4096)
	go ClientRead(conn, buffer, exit)
	go ClientSend(conn, exit)
	t := time.NewTicker(100 * time.Second)
	for {
		select {
		case <-exit:
			goto Exit
		case <-t.C:
			fmt.Println("当前启动线程数:", runtime.NumGoroutine())
		}
	}
Exit:
	fmt.Println("客户端关闭")

}
