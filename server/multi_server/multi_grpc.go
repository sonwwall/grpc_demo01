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
	proto.UnimplementedVideoServiceServer
}

type OrderServer struct {
	proto.UnimplementedOrderServiceServer
}

func (VideoServer) Look(ctx context.Context, request *proto.Request) (res *proto.Response, err error) {
	fmt.Println("video:", request)
	res = &proto.Response{
		Name: "外城",
	}
	return res, nil
}

func (OrderServer) Buy(ctx context.Context, request *proto.Request) (res *proto.Response, err error) {
	fmt.Println("order:", request)
	res = &proto.Response{
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

	proto.RegisterVideoServiceServer(s, &VideoServer{})
	proto.RegisterOrderServiceServer(s, &OrderServer{})
	fmt.Println("服务端正在运行在8080端口")
	err = s.Serve(listen)
	if err != nil {
		grpclog.Fatal("serve err:", err.Error())
	}
}
