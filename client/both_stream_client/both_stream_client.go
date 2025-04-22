package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"grpc_demo01/stream_proto/stream"
)

func main() {
	addr := ":8080"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		grpclog.Fatal("connection err:", err.Error())
	}
	defer conn.Close()

	client := stream.NewBothStreamClient(conn)
	streaming, err := client.Chat(context.Background())

	for i := 0; i < 10; i++ {
		streaming.Send(&stream.Request{Name: fmt.Sprintln("你好呀，这是我第", i+1, "次给你发消息")})
		response, err := streaming.Recv()
		fmt.Println(response, err)

	}
}
