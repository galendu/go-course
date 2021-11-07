package main

import (
	"context"
	"fmt"

	"gitee.com/infraboard/go-course/day21/grpc01/service"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := service.NewHelloServiceClient(conn)
	resp, err := client.Hello(context.Background(), &service.Request{Value: "hi grpc server"})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
