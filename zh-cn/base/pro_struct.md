# 程序结构

## GOPATH设置

go path 详细用途说明: 

```sh
go help gopath
```

读取环境变量的配置:

```sh go env <NAME>
go env 
```

查看环境变量的含义:
```sh
go help environment
```

更改配置: go env -w <NAME>=<VALUE>

```sh
go env -w GOPATH=<workspace>
```

以上 $GOPATH 目录约定有三个子目录：
+ src 存放源代码（比如：.go .c .h .s等）
+ pkg 编译后生成的静态库（比如：.a）, 多个静态库文件通过连接器连接 最终构成我们要得目标文件
+ bin 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录）

## 第一个Go程序

```go
package main

import "fmt"

func main() {
    /* 这是我的第一个简单程序 */
    fmt.Println("Hello, World!")
}
```

解读:
+ package: Go源文件开头必须使用 package 声明代码所属包，包是 Go 代码分发的最基本单位。若程序需要运行包名必须为 main。
+ import: 用于导入程序依赖的所有的包。此程序依赖于 fmt 包。
+ func: 用于定义函数。main 函数是程序的入口,若程序需要运行必须声明 main 函数,main
函数无参数也无返回值
+ fmt.Println 调用 fmt.Println 函数将参数信息打印到控制台

Go语言的基础组成有以下几个部分：
 + 包声明
 + 引入包
 + 函数
 + 变量
 + 语句 & 表达式
 + 注释

## 运行程序
1. go build: 用于编译&链接程序或包
2. go build -work -x -o helloworld.exe main.go
3. go run：用于直接运行程序
4. go run -work -x main.go
5. go clean：清除编译文件
6. 常用参数：
    + -x: 打印编译过程执行的命令，并完成编译或运行
    + -n: 只打印编译过程执行命令
    + -work：打印编译过程的临时目录
    + -o: 指定编译结果文件


## Go 工具链介绍

### go install 与 go tool compile/link

1. 关闭go mod

```sh
go env -w GO111MODULE=off
```

2. 编译pkg

```sh
go install ./pkg/
```

3. 编译main.go

```sh
go tool compile -I /e/Golang/pkg/windows_amd64 main.go
```

4. 链接main.o

```sh
go tool link -o main.exe -L /e/Golang/pkg/windows_amd64/  main.o
```

## go clean

这个命令是用来移除当前源码包和关联源码包里面编译生成的文件

```sh
ls /e/Golang/pkg/windows_amd64/demo
```

```sh
go clean -i -n
```

记得恢复mod模式
```sh
go env -w GO111MODULE=on
```

## go fmt

代码格式化


## go get

下载源码包并执行go install

## go vet

语法静态检查

## 作业

+ 读完所有go命令行工具的说明文档, 理解所有go env环境变量的含义
+ 理解Go Mod和Go Path的作用机制
+ 使用go tool 工具链自己编译程序, 理解编译和链接的含义