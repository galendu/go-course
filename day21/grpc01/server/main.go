package main

import (
	"context"
	"net"

	"gitee.com/infraboard/go-course/day21/grpc01/service"
	"google.golang.org/grpc"
)

// 通过接口约束HelloService服务
var _ service.HelloServiceServer = (*HelloService)(nil)

type HelloService struct {
	service.UnimplementedHelloServiceServer
}

func (p *HelloService) Hello(ctx context.Context, req *service.Request) (*service.Response, error) {
	resp := &service.Response{}
	resp.Value = "hello:" + req.GetValue()
	return nil, nil
}

func main() {
	server := grpc.NewServer()
	service.RegisterHelloServiceServer(server, &HelloService{})

	lis, err := net.Listen("tcp", "1234")
	if err != nil {
		panic(err)
	}

	server.Serve(lis)
}
