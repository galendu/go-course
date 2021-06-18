package day3

import (
	"fmt"
	"testing"
)

func TestPanic1(t *testing.T) {
	fn()
}

func fn() {
	fmt.Println("start fn")
	panic("pannic in fn")
	fmt.Println("end fn")
}

func TestPanic2(t *testing.T) {
	var a *int
	fmt.Println(*a)
}

func TestPanic3(t *testing.T) {
	var x, y *int
	sum(x, y)
}

func sum(x, y *int) int {
	return *x + *y
}

func TestPanic4(t *testing.T) {
	var x, y *int
	fmt.Println(recover())
	sum(x, y)
}

func TestPanic5(t *testing.T) {
	defer func() {
		fmt.Println(recover())
	}()

	var x, y *int
	sum(x, y)
}
