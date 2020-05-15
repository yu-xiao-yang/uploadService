package  main

import "individual/UploadService/service"

func main() {
	server:=service.UploadServe{}
	server.Server()
}
