package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func runTask(id int) {
	// 推出一个减去1
	defer wg.Done()

	fmt.Printf("task %d start..\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("task %d complete\n", id)
}

func asyncRun() {
	for i := 0; i < 10; i++ {
		go runTask(i + 1)
		// 没启动一个go routine 就+1
		wg.Add(1)
	}
}

func main() {
	asyncRun()
	wg.Wait()
}
