# Demo后端

功能: CMDB主机信息录入与查询

涉及到的技能:

+ go http标准库
+ 第三方路由库: [httprouter](https://github.com/julienschmidt/httprouter)
+ go 操作mysql


## 项目骨架介绍


## 数据结构与接口定义


## 基于MySQL存储的接口实现


## HTTP API暴露

我们以RestFull风格来设计我们的接口

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func QueryUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "query user!\n")
}

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "create user!\n")
}

func DescribeUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "describe, %s!\n", ps.ByName("name"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "delete, %s!\n", ps.ByName("name"))
}

var (
	addr = ":8080"
)

func main() {
	router := httprouter.New()
	router.GET("/users", QueryUser)
	router.POST("/users", CreateUser)
	router.GET("/users/:id", DescribeUser)
	router.DELETE("/users/:id", DeleteUser)

	log.Printf("listen on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
```

## 如何管理项目配置



