package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"grpc_demo01/stream_proto/stream"
	"io"
	"log"
	"os"
)

func main() {
	addr := ":8080"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		grpclog.Fatal("connection err:", err.Error())
	}
	defer conn.Close()

	client := stream.NewClientStreamClient(conn)
	streaming, err := client.UploadFile(context.Background())

	file, err := os.Open("static/01.jpg")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()
	for {
		buf := make([]byte, 1024)
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to read file: %v", err)
		}
		streaming.Send(&stream.FileRequest{Content: buf, FileName: "02.jpg"})
	}

	//for i := 0; i < 10; i++ {
	//	err := streaming.Send(&stream.FileRequest{FileName: fmt.Sprintf("file%d", i)})
	//
	//	if err == io.EOF {
	//		fmt.Println("结束")
	//	}
	//	if err != nil {
	//		fmt.Println(err.Error())
	//	}
	//}
	response, err := streaming.CloseAndRecv()
	fmt.Println(response, err)

}
