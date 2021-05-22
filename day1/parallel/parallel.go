package main

import (
	"fmt"
	"runtime"
	"time"

	"golang.org/x/sys/windows"
)

func main() {
	changeMax(6)
	startG(12)
	select {}
}

func startG(n int) {
	for i := 0; i < n; i++ {
		go task(false)
	}
}

func changeMax(n int) {
	pre := runtime.GOMAXPROCS(n)
	fmt.Printf("previous: %d, set: %d, total: %d\n", pre, n, runtime.NumCPU())
}

func task(lockThread bool) {
	if lockThread {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}

	createID := windows.GetCurrentThreadId()
	fmt.Printf("thread id: %d\n", createID)
	for {
		time.Sleep(2 * time.Second)
		fmt.Println(createID, windows.GetCurrentThreadId())
	}
}
