# prometheus定制开发

我们并不会修改prometheus的源码




## 概念简介








## 开发环境搭建







## Exporter 开发


### 简单粗暴

我们之间开发一个满足prometheus格式的API接口即可

```go
package main

import (
    "fmt"
    "net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "lexporter_request_count{user=\"admin\"} 1000" )
}

func main () {
    http.HandleFunc("/metrics", HelloHandler)
    http.ListenAndServe(":8000", nil)
}
```



### 使用SDK

```

```