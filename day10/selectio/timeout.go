package selectio

import (
	"fmt"
	"time"
)

func SelectTimeout() {
	ch1 := make(chan string)

	// 激活一个goroutine，但5秒之后才发送数据
	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "put value into ch1"
	}()

	select {
	case val := <-ch1:
		fmt.Println("recv value from ch1:", val)
		return
	// 只等待3秒，然后就结束
	case <-time.After(3 * time.Second):
		fmt.Println("3 second over, timeover")
	}
}
