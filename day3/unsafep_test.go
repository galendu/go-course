package day3

import (
	"fmt"
	"testing"
)

func TestSafePointer1(t *testing.T) {
	str := "pointer_test"
	a := &str
	fmt.Println(a)
}
