package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"grpc_demo01/grpc_proto/hello_grpc"
	"net"
)

// 这个结构体名字不重要
type HelloService struct {
	hello_grpc.UnimplementedHelloServiceServer
}

type HelloWorldService struct {
	hello_grpc.UnimplementedHelloWorldServer
}

func (HelloService) SayHello(ctx context.Context, request *hello_grpc.HelloRequest) (res *hello_grpc.HelloResponse, err error) {
	fmt.Println(request)
	return &hello_grpc.HelloResponse{
		Name:    "外城",
		Message: "ok",
	}, nil
}

func (HelloWorldService) SayHelloWorld(ctx context.Context, request *hello_grpc.HelloWorldRequest) (res *hello_grpc.HelloWorldResponse, err error) {
	fmt.Println(request)
	name := request.Name
	fmt.Println("HelloWorld")
	return &hello_grpc.HelloWorldResponse{Name: name}, nil
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 创建一个gRPC服务器实例。
	s := grpc.NewServer()
	server := HelloService{}
	server02 := HelloWorldService{}
	// 将server结构体注册为gRPC服务。
	hello_grpc.RegisterHelloServiceServer(s, &server)
	hello_grpc.RegisterHelloWorldServer(s, &server02)
	fmt.Println("grpc server running :8080")
	// 开始处理客户端请求。
	err = s.Serve(listen)
}
