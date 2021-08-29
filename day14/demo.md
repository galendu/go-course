# 前端后端项目样例

功能: 通过网页查看用户列表

## API 开发

涉及到的技能:

+ go http标准库
+ 第三方路由库: [httprouter](https://github.com/julienschmidt/httprouter)
+ go 操作mysql


### HTTP 接口编写

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


## UI 开发

