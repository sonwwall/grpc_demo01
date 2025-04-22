package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"grpc_demo01/grpc_proto/multi_grpc"
	"net"
)

type VideoServer struct {
	multi_grpc.UnimplementedVideoServiceServer
}

type OrderServer struct {
	multi_grpc.UnimplementedOrderServiceServer
}

func (VideoServer) Look(ctx context.Context, request *multi_grpc.Request) (res *multi_grpc.Response, err error) {
	fmt.Println("video:", request)
	res = &multi_grpc.Response{
		Name: "外城",
	}
	return res, nil
}

func (OrderServer) Buy(ctx context.Context, request *multi_grpc.Request) (res *multi_grpc.Response, err error) {
	fmt.Println("order:", request)
	res = &multi_grpc.Response{
		Name: "苹果",
	}
	return res, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		grpclog.Fatal("listen err:", err.Error())
	}

	s := grpc.NewServer()

	multi_grpc.RegisterVideoServiceServer(s, &VideoServer{})
	multi_grpc.RegisterOrderServiceServer(s, &OrderServer{})
	fmt.Println("服务端正在运行在8080端口")
	err = s.Serve(listen)
	if err != nil {
		grpclog.Fatal("serve err:", err.Error())
	}
}
