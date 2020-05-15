package service

import (
	"fmt"
	"io"
	"os"
	"path"
)

//文件头
type FileUploadHead struct{
	filename string  `json:"filename"`//文件名称
	fileSize int64  `json:"file_size"`//文件大小
	sliceInterval int64  `json:"slice_interval"`//分片间隔
	checksum string `json:"checksum"` //完整校验码
}

func fileSlice(filelocation string, interval  int64) error{

	_,filename := path.Split(filelocation)
	fmt.Println(filename)
	if finfo,err:=os.Lstat(filelocation );err==nil{
		fmt.Println("name: ",finfo.Name(),"  size:",finfo.Size()," mode:",finfo.Mode().String())
	}

	wfile,err:=os.Create("/home/yuyang/test/test.go")
	defer wfile.Close()
	if err!=nil {
		return  err
	}
	//md5.New().Sum()
	var index int64 = 0
	fileObj,err :=os.Open(filelocation);
	if err!=nil {
		return  err
	}
	defer fileObj.Close()

	for{
			data:=make([]byte,1500)
			lendata,err:=fileObj.ReadAt( data,interval*index)
			if err!=nil {
				fmt.Println(err)
			}
			if err!=nil && err!=io.EOF {
				return err
			}
			if _,err:=wfile.WriteAt(data,index*interval); err!=nil{
				return err
			}

			if(err==io.EOF || lendata ==0 ) {
				break
			}
			index ++
			fmt.Println("index:",index," len:", lendata)
	}

	return  nil
}