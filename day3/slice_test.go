package day3

import (
	"fmt"
	"testing"
)

func TestSliceSize(t *testing.T) {
	a := make([]int, 3, 5)
	fmt.Println(a, len(a), cap(a))
	fmt.Println(a[2])
	fmt.Println(a, len(a), cap(a))
}
