package main

import (
	"fmt"
	"log"
	"net"
	"runtime"
)

//var html = `<!DOCTYPE html>
//<html lang="en">
//<head>
//    <meta charset="UTF-8">
//    <title>magedu</title>
//</head>
//<body>
//<h1>马哥教育www.magedu.com -- Goroutine</h1>
//</body>
//</html>`
//
//var head = `HTTP/1.1 200 OK
//Date: Mon, 24 Oct 2022 20:04:23 GMT
//Content-Type: text/html
//Content-Length: %d
//Connection: keep-alive
//Server: wayne.magedu.com
//
//%s`
//
//var response = fmt.Sprintf(head, len(html), html)

var response = "true"

func Accept(s *net.TCPListener, b []byte) {
	conn, err := s.Accept() //接收连接，分配socket
	if err == nil {
		log.Printf("客户端%s连接成功", conn.RemoteAddr().String())
	} else {
		log.Panicln(err)
	}

	go Con(conn, b)

}

func Con(conn net.Conn, b []byte) {
	defer conn.Close()
	buffer := b
	for {
		n, err := conn.Read(buffer)
		fmt.Printf("从客户端%v接收到数据%v\n", conn.RemoteAddr().String(), string(b[:n]))
		if n == 0 {
			log.Printf("客户端%s主动断开", conn.RemoteAddr().String())
			return
		}
		if err != nil {
			log.Panicln(err)
			return
		}
		conn.Write([]byte(response))
	}
}
func GetAddr() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Panicln(err)
	}
	var s string
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				s = fmt.Sprint(ipnet.IP.String())
				break
			}
		}
	}
	return s
}

func main() {
	var numCPU = runtime.NumCPU()
	fmt.Printf("当前CPU核心数%d\n", numCPU)
	a := GetAddr()
	addr := fmt.Sprint(a + ":9999")
	fmt.Printf("在%s上启动服务\n", addr)
	laddr, err := net.ResolveTCPAddr("tcp4", addr) //解析地址
	if err != nil {
		log.Panicln(err) //panicln会打印异常，程序退出
	}
	server, err := net.ListenTCP("tcp4", laddr)
	if err != nil {
		log.Panicln(err)
	}
	defer server.Close() //用完关闭
	for {
		buffer := make([]byte, 4096)
		Accept(server, buffer)
	}
}
