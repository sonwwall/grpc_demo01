package main

import (
	"fmt"
	"net/rpc"
)

type Req struct {
	Num1 int
	Num2 int
}
type Res struct {
	Num int
}

func main() {
	req := Req{1, 2}
	client, err := rpc.DialHTTP("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	var res Res
	err = client.Call("Server.Add", req, &res) //调用Server的Add方法
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}
