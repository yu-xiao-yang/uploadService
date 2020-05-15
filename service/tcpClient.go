package service

import (
	"fmt"
	"net"
	"time"
)


func Client() {
	fmt.Println("client launch")
	serverAddr := "localhost:8888"
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		fmt.Println("Resolve TCPAddr error", err)
	}
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println("connect server error", err)
	}

	conn.Write([]byte("hello , I am client"))
	go recv(conn)
	time.Sleep(2 * time.Second)//等两秒钟，不然还没接收数据，程序就结束了。
}

func recv(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err == nil {
		fmt.Println("read message from server:" + string(buffer[:n]))
		fmt.Println("Message len:", n)
	}
}
