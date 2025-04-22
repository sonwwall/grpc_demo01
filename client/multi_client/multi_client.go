package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_demo01/grpc_proto/multi_grpc"
	"log"
)

func main() {
	addr := ":8080"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	orderClient := multi_grpc.NewOrderServiceClient(conn)
	videoClient := multi_grpc.NewVideoServiceClient(conn)
	res, err := orderClient.Buy(context.Background(), &multi_grpc.Request{Name: "龙哥"})
	fmt.Println(res)
	res, err = videoClient.Look(context.Background(), &multi_grpc.Request{Name: "龙哥"})
	fmt.Println(res)

}
