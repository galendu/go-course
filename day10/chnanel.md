## Channel基础

我们除了可以通过共享内存进行goroutine之间的通信, Go还提供一个特殊的数据类型也可以用于多个goroutine之间的通信: chan

在Go的并发编程范式里面是不主张使用共享内存的方式来通信，

### 创建channel

channel是指针类型的数据类型，通过make来分配内存。例如

```go
ch := make(chan int)
```

这表示创建一个channel，这个channel中只能保存int类型的数据。也就是说一端只能向此channel中放进int类型的值，另一端只能从此channel中读出int类型的值。

需要注意，chan TYPE才表示channel的类型。所以其作为参数或返回值时，需指定为xxx chan int类似的格式


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