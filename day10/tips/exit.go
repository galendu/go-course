package tips

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(cannel chan struct{}) {
	for {
		select {
		default:
			fmt.Println("hello")
			time.Sleep(100 * time.Millisecond)
		case <-cannel:
			// 退出
		}
	}
}

func CancelWithChannel() {
	cannel := make(chan struct{})
	go worker(cannel)

	time.Sleep(time.Second)
	cannel <- struct{}{}
}

func workerv2(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
			time.Sleep(100 * time.Millisecond)
		case <-cancel:
			return
		}
	}
}

func CancelWithDown() {
	cancel := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerv2(&wg, cancel)
	}

	time.Sleep(time.Second)

	// 发送退出信号
	close(cancel)

	// 等待goroutine 安全退出
	wg.Wait()
}

func workerV3(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
			time.Sleep(100 * time.Millisecond)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
func CancelWithCtx() {
	// 控制超时
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go workerV3(ctx, &wg)
	}

	time.Sleep(time.Second)

	// 取出任务
	cancel()

	// 等待安全退出
	wg.Wait()
}
