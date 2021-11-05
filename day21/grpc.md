# GRPC 入门

![](./images/grpc.png)

gRPC是Google公司基于Protobuf开发的跨语言的开源RPC框架。gRPC基于HTTP/2协议设计，可以基于一个HTTP/2链接提供多个服务，对于移动设备更加友好。本节将讲述gRPC的简单用法

这个是我们之前定义的接口
```go
package service

const HelloServiceName = "HelloService"

type HelloService interface {
	Hello(*Request, *Response) error
}
```

其中数据结构: Request和Response 已经使用protobuf定义了数据的交换格式,  如果我们的接口也能通过 protobuf定义是不是 就完美了, 这也是GRPC真正的威力


## GRPC技术栈

![](./images/grpc-go-stack.png)

+ 数据交互格式: protobuf
+ 通信方式: 最底层为TCP或Unix Socket协议，在此之上是HTTP/2协议的实现
+ 核心库: 在HTTP/2协议之上又构建了针对Go语言的gRPC核心库
+ Stub: 应用程序通过gRPC插件生产的Stub代码和gRPC核心库通信，也可以直接和gRPC核心库通信

gRPC采用protobuf描述 接口和数据, 我们可以把他理解为: protobuf ON HTTP2 的一种RPC


## Hello gRPC

下面我们讲演示一个基础的gRPC服务.

### protobuf grpc插件

protobuf 不仅可以定义交互的数据结构(message), 还可以定义交互的接口:

```protobuf
service HelloService {
    rpc Hello (String) returns (String);
}
```

从Protobuf的角度看，gRPC只不过是一个针对service接口生成代码的生成器。因此我们需要提前安装grpc的代码生成插件

```sh
# protoc-gen-go 插件之前已经安装
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 安装protoc-gen-go-grpc插件
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

我们看一看到当前插件的版本
```sh
protoc-gen-go-grpc --version                                   
protoc-gen-go-grpc 1.1.0
```

### 生成代码

然后基于protoc-gen-go-grpc来生产我们的grpc代码, 我们把之前的rpc 修改为GRPC

我们看看protobuf 定义接口的语法:
```protobuf
service <service_name> {
    rpc <function_name> (<request>) returns (<response>);
}
```
+ service: 用于申明这是个服务的接口
+ service_name: 服务的名称,接口名称
+ function_name: 函数的名称
+ request: 函数参数， 必须的
+ response: 函数返回， 必须的, 不能没有

```protobuf
syntax = "proto3";

package hello;
option go_package="gitee.com/infraboard/go-course/day21/pbrpc/service";

// The HelloService service definition.
service HelloService {
    rpc Hello (Request) returns (Response);
}

message Request {
    string value = 1;
}

message Response {
    string value = 1;
}
```

然后我们生产代码, 同时制定gprc插件对应参数:
```sh
protoc -I=. --go_out=./grpc/service --go_opt=module="gitee.com/infraboard/go-course/day21/grpc/service" \
--go-grpc_out=./grpc/service --go-grpc_opt=module="gitee.com/infraboard/go-course/day21/grpc/service" \
grpc/service/service.proto
```

生成的客户端: 
```go
// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}
```

生成的服务端:
```go
// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	Hello(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedHelloServiceServer()
}
```


## 参考

+ [GRPC Quick start](https://grpc.io/docs/languages/go/quickstart/)
+ [GRPC Examples](https://github.com/grpc/grpc-go/tree/master/examples)