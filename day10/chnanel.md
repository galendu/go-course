# CSP与Chnanel

我们除了可以通过共享内存进行goroutine之间的通信, Go还提供一个特殊的数据类型也可以用于多个goroutine之间通信: chan

虽然我们在 Go 语言中也能使用共享内存加互斥锁进行通信，但是 Go 语言提供了一种不同的并发模型，即通信顺序进程（Communicating sequential processes，CSP）。

Goroutine 和 Channel 分别对应 CSP 中的实体和传递信息的媒介，Goroutine 之间会通过 Channel 传递数据

![](../image/go-chan.jpg)

## Channel基础

channel是指针类型的数据类型，使用 chan表示, 通过make来分配内存。例如

```go
ch := make(chan int)
```

这表示创建一个channel，这个channel中只能保存int类型的数据。也就是说一端只能向此channel中放进int类型的值，另一端只能从此channel中读出int类型的值。

需要注意，chan TYPE才表示channel的类型。所以其作为参数或返回值时，需指定为xxx chan int类似的格式

1.往channel中发送消息
```go
ch <- VALUE
```

2.从channel中获取消息
```go
<-ch               // 取出消息，直接扔掉
value := <-ch      // 从ch中读取一个值并保存到value变量中
value,ok = <-ch    // 从ch读取一个值，判断是否读取成功，如果成功则保存到value变量中
for v := range ch  // 通过for range语法来取值, 知道ch关闭时退出循环
```

简单总结为:
+ send: 当ch出现在<-的左边
+ recv: 当ch出现在<-的右边

比如: alice -> bob 发送一个消息(hello, this is alice)
```go
package main

import (
	"fmt"
	"time"
)

func sender(ch chan string) {
	ch <- "hello"
	ch <- "this"
	ch <- "is"
	ch <- "alice"
}

func recver(ch chan string) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan string)
	go sender(ch) // sender goroutine
	go recver(ch) // recver goroutine

	time.Sleep(1 * time.Second)
}
```




### channel的操作

每个channel都有3种操作：send、receive和close
+ send：表示sender端的goroutine向channel中投放数据, 格式如下: ch <- VALUE
+ receive：表示receiver端的goroutine从channel中读取数据, 从ch中读取一个值 val,ok := <-ch          
+ close：表示关闭channel, 比如 close(ch)

注意事项:
+ 关闭channel后，send操作将导致painc
+ 关闭channel后，recv操作将返回对应类型的0值以及一个状态码false
+ close并非强制需要使用close(ch)来关闭channel，在某些时候可以自动被关闭
+ 如果使用close()，建议条件允许的情况下加上defer, 只在sender端上显式使用close()关闭channel。因为关闭通道意味着没有数据再需要发送


### 实战: AB交互打印

```go
package main

import (
	"fmt"
	"time"
)

func A(startA, startB chan struct{}) {
	a := []string{"1", "2", "3"}
	index := 0
	for range startA {
		if index > 2 {
			return
		}
		fmt.Println(a[index])
		index++
		startB <- struct{}{}
	}
}

func B(startA, startB chan struct{}) {
	b := []string{"x", "y", "z"}
	index := 0
	for range startB {
		fmt.Println(b[index])
		index++
		startA <- struct{}{}
	}
}

func main() {
	startA, startB := make(chan struct{}), make(chan struct{})
	go A(startA, startB)
	go B(startA, startB)

	startA <- struct{}{}
	time.Sleep(1 * time.Second)
}
```