# prometheus定制开发

我们并不会修改prometheus的源码




## 概念简介








## 开发环境搭建







## Exporter 开发


### 简单粗暴

我们直接开发一个满足prometheus格式的API接口即可

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

并不是所有的代码都需要我们自己去实现, Prometheus为我们准备了一个客户端, 我们可以基于客户端快速添加监控

```go
package main

import (
 "net/http"

 "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    // Serve the default Prometheus metrics registry over HTTP on /metrics.
 http.Handle("/metrics", promhttp.Handler())
 http.ListenAndServe(":8080", nil)
}
```

### 自定义指标




### 指标类型
