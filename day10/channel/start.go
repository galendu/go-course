package chnanel

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

func SyncAB() {
	startA, startB := make(chan struct{}), make(chan struct{})
	go A(startA, startB)
	go B(startA, startB)

	startA <- struct{}{}
	time.Sleep(1 * time.Second)
}
