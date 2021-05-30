# 函数

### 内置函数
在 buildin/buildin.go内定义了Go所有支持内置函数：make、len、cap、new、append、copy、close、delete、complex、real、 imag、panic、recover

我们在代码里面可以直接使用: 比如使用len计算字符串的长度
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(len("asdfsdf"))
}
```

内置函数和用户定义函数没有本质上的区别, 我们也可以覆盖它
```go
package main

import (
	"fmt"
)

func len(a string) string {
	return "hello"
}

func main() {
	fmt.Println(len("asdfsdf"))
}
```
但是你最好不要这样做，你会被打的