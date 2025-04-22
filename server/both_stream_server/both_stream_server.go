package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_demo01/stream_proto/stream"
	"log"
	"net"
)

type Server struct {
	stream.UnimplementedBothStreamServer
}

func (Server) Chat(streaming stream.BothStream_ChatServer) error {

	for i := 1; i < 10; i++ {
		request, _ := streaming.Recv()
		fmt.Println(request)
		streaming.Send(&stream.Response{Name: "我已经收到了你的消息"})
	}

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	stream.RegisterBothStreamServer(server, &Server{})

	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
