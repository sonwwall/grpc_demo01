package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_demo01/service_proto/proto"
	"log"
)

func main() {
	addr := ":8080"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	orderClient := proto.NewOrderClient(conn)
	videoClient := proto.NewVideoClient(conn)
	res, err := orderClient.Buy(context.Background(), &proto.Request{Name: "龙哥"})
	fmt.Println(res)
	res, err = videoClient.Look(context.Background(), &proto.Request{Name: "龙哥"})
	fmt.Println(res)

}
