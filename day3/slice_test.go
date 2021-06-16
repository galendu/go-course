package day3

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSliceSize(t *testing.T) {
	a := make([]int, 3, 5)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(a[2])
	fmt.Println(a, len(a), cap(a))
}

func TestSliceAddr(t *testing.T) {
	a := make([]int, 3, 5)
	fmt.Printf("%p\t", a)
	fmt.Println(&(a[0]), &(a[1]), 0xa3c0-0xa3c8)

	b := make([]int64, 3, 5)
	fmt.Println(unsafe.Sizeof(&(b[0])))
	fmt.Println(&(b[0]), &(b[1]), 0xa3c0-0xa3c8)
}
