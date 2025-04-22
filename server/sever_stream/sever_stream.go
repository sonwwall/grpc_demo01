package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo01/stream_proto/stream"
	"io"
	"log"
	"net"
	"os"
)

type ServiceStream struct {
	stream.UnimplementedServiceStreamServer
}

func (ServiceStream) Fun(requset *stream.Request, streaming stream.ServiceStream_FunServer) error {
	fmt.Println(requset)
	for i := 1; i < 10; i++ {
		err := streaming.Send(&stream.Response{
			Name: fmt.Sprintln("这是我第", i+1, "次响应"),
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
	return nil
}

func (ServiceStream) DownLoadFile(requset *stream.Request, streaming stream.ServiceStream_DownLoadFileServer) error {
	fmt.Println(requset)

	file, err := os.Open("static/高数下册-2022版讲义.pdf")
	if err != nil {
		return err
	}
	defer file.Close()
	for {
		buf := make([]byte, 1024)
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		streaming.Send(&stream.FileResponse{Content: buf})
	}

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	stream.RegisterServiceStreamServer(server, &ServiceStream{})

	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
