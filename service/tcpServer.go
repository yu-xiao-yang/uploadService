package service

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type UploadServe struct {
	IsStop chan bool
	mconns map[string]*net.Conn//connect handles
}
func NewUploadServer() (*UploadServe){
	return &UploadServe{
		IsStop:make(chan bool,2),
		mconns:make(map[string]*net.Conn ),
	}
}

func getkey(ip net.IP, port uint16) string {
	return ip.String()+"_"+strconv.Itoa(int(port) )
}
func (this UploadServe) RemoteConn(key string){
	delete(this.mconns,key)
}
func  (this UploadServe) SetConn(key string,val *net.Conn){
	this.mconns[key] = val
}
func (this UploadServe) GetConn(key string) (*net.Conn,bool){
	val,ok :=this.mconns[key]
	return  val ,ok
}
func (this UploadServe) Stop() {
	this.IsStop <- true
}
func (this UploadServe)Wait(){
	select {
		case <-this.IsStop :
			//TODO REALSE ALL

	}
}
func (this UploadServe) handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("connection success")
	fmt.Println("client address: ", conn.RemoteAddr())
	this.SetConn(conn.RemoteAddr().String() ,&conn)

	buffer := make([]byte, 1024)
	recvLen, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Read error", err)
	}
	strBuffer := string(buffer[:recvLen])
	fmt.Println("Message: ", strBuffer)
	fmt.Println("Message len :", recvLen)
	time.Sleep(time.Second * 1)//等一秒钟，可以看出client里面的read函数有阻塞效果


	sendLen, err := conn.Write([]byte("I am server, you message :" + strBuffer))//将client发过来的消息原样发送回去
	if err != nil {
		fmt.Println("send message error", err)
	}
	fmt.Println("send message success")
	fmt.Println("send message len；", sendLen)
}

func (this UploadServe) Server() {
	fmt.Println("hello world")
	lner, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("listener creat error", err)
	}
	fmt.Println("waiting for client")
	for {
		conn, err := lner.Accept()
		if err != nil {
			fmt.Println("accept error", err)
		}
		go this.handleConnection(conn)
	}

}
