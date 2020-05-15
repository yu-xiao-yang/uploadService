package service

import (
	"fmt"
	"net"
)

func recieveData(conn net.Conn){
	 buf:=make([]byte, 1500)
	 if len,err:=conn.Read(buf);err==nil{
		 fmt.Println(string(buf[:len]))
		 fmt.Print("len:",len)
	 }
}
