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
