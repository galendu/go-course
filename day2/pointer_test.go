package day2

import (
	"fmt"
	"testing"
)

func TestPointerRaw(t *testing.T) {
	a := "string a"
	fmt.Println(&a)
}

func TestPointer1(t *testing.T) {
	var a *int
	fmt.Println(a)
}

func TestPointer2(t *testing.T) {
	var a *int = new(int)
	fmt.Println(a)
	fmt.Println(*a)
}
