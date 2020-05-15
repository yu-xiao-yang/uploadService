package service

import (
	"fmt"
	"testing"
)

func TestFileprocess(t *testing.T){
	if err:=fileSlice("/home/yuyang/src/gopro/src/individual/UploadService/service/fileProcess.go", 1024);err!=nil{
		fmt.Println("err:",err)
	}
}

func TestParsePayload(t *testing.T){
	mh:=MessageHead{}
	mh.EventType = uint16(1)
	mh.Headlength = uint16(16)
	mh.MsgType = uint16(1)
	mh.Version = uint16(1)
	data,err:=Encode(mh)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println("lenbyte:",len(data))
	m2:=&MessageHead{}
	Decode(data,m2)
	fmt.Println(m2.toString())
	//ParsePayload()
}