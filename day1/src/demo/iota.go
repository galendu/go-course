package main

import (
	"demo/pkg"
	"fmt"
)

const (
	MALE   string = "MAN"
	FEMALE string = "WOMEN"
)

const (
	a = iota + 1 // a = 0
	b            // b = a + 1
	c            // c =b + 1
)

var (
	DefaultName = "laoyu"
)

type Hello interface {
	Hello(username string) string
}

type CoustB func(x, y int) int

func PointerTest(p *string) string {
	*p = "string b"
	return *p
}

// Marka
// Markb
func main() {
	fmt.Println(a, b, c)

	CoustB := func(x, y int) int { return 1 }
	fmt.Println(CoustB)

	aStr := "string a"
	fmt.Println(PointerTest(&aStr))
	fmt.Println(a)
	pkg.Demo()
	fmt.Print("Hello world")
}
