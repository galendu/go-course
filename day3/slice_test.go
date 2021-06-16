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

func TestSlice(t *testing.T) {
	hlist := make(map[string][]string)

	h1key := []string{}
	h1key = append(h1key, "1", "2", "3")
	hlist["1"] = h1key

	fmt.Println(hlist)
}
