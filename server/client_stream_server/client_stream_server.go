package main

import (
	"bufio"
	"google.golang.org/grpc"
	"grpc_demo01/stream_proto/stream"
	"io"
	"log"
	"net"
	"os"
)

type ClientStream struct {
	stream.UnimplementedClientStreamServer
}

func (ClientStream) UploadFile(streaming stream.ClientStream_UploadFileServer) error {

	file, err := os.OpenFile("static/02.jpg", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for {
		response, err := streaming.Recv()
		if err == io.EOF {
			break
		}
		writer.Write(response.Content)
	}

	//for i := 1; i < 10; i++ {
	//	response, err := streaming.Recv()
	//
	//	fmt.Println(response, err)
	//}
	streaming.SendAndClose(&stream.Response{Name: "完成"})

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	stream.RegisterClientStreamServer(server, &ClientStream{})

	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
