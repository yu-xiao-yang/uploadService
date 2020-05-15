package service

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"


	"unsafe"
)

const BUFFER_MAX_LEN  =1500


//event type
const (
	UPLOAD_FILE_START = 1
	UPLOAD_FILE_END = 2
	UPLOAD_FILE_RESTART  =3
)

//message type
const (
	FILE_BINARY_DATA = 1
	FILE_HEAD_DESCRIBE =2
)

/*
|-------------------------------------------|
|    file head                              |
|-------------------------------------------|
| segment1 head + segment body              |
|___________________________________________|
|.....segment2 head +segment body           |
|___________________________________________|
*/


//消息头
type MessageHead struct {
	Version uint16 // 消息版本
	Headlength uint16 //消息长度
	EventType  uint16  //0:客户端  1:server 端
	MsgType uint16  // 消息类型
}
func (this *MessageHead) toString() string{
	return  fmt.Sprintf(" version:%d headlength:%d,eventType:%d, msgType:%d",this.Version,
		this.Headlength,this.EventType,this.MsgType)
}
//消息体
type MessageContent struct {
	MessageHead
	body []byte  //数据
}
func  Encode(obj interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, obj); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type  DecodeCallback func(obj interface{})

func Decode(b []byte,obj *MessageHead) ( error) {
	buf := bytes.NewBuffer(b)
	if err := binary.Read(buf, binary.LittleEndian, obj); err != nil {
		return  err
	}
	return  nil
}

func ParsePayload(payload []byte ) error{
	length:= len(payload)
	headlen := 4
	//headlen,err:=struc.Sizeof(MessageHead{})
	fmt.Println(length  ,"> ", headlen)
	//if err!=nil {
	//	return  err
	//}
	if length < headlen{
		return  errors.New("head length  error")
	}
	msgh := (*MessageHead)(unsafe.Pointer(&payload) )
	fmt.Println(msgh.toString() )
	return  nil
}



