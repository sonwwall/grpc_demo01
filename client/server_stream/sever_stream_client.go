package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"grpc_demo01/stream_proto/stream"
	"io"
	"os"
)

func main() {
	addr := ":8080"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		grpclog.Fatal("connection err:", err.Error())
	}
	defer conn.Close()

	client := stream.NewServiceStreamClient(conn)
	//streaming, err := client.Fun(context.Background(), &stream.Request{Name: "你好"})
	//if err != nil {
	//	grpclog.Fatal("接收消息失败:", err.Error())
	//}
	//for {
	//	response, err := streaming.Recv()
	//	if err == io.EOF {
	//		return
	//	}
	//	fmt.Println(response, err)
	//}

	//下载文件
	streaming, err := client.DownLoadFile(context.Background(), &stream.Request{Name: "请求一次"})
	file, err := os.OpenFile("static/高数下册-2022版讲义1.pdf", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for {
		response, err := streaming.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}
		fmt.Println("写入", len(response.Content), "数据")
		writer.Write(response.Content)
	}
	writer.Flush()
}
