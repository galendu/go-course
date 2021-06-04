package day2

import (
	"fmt"
	"testing"
)

func TestForBase(t *testing.T) {
	var sum int
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func TestForShort(t *testing.T) {
	var sum int
	for sum < 10 {
		sum++
	}

	fmt.Println(sum)
}

func TestForLoop(t *testing.T) {
	var sum int
	for {
		sum++
		fmt.Println(sum)
		if sum == 100 {
			return
		}
	}
}

func TestForRange(t *testing.T) {
	iter := "abcdefg"
	for index, value := range iter {
		fmt.Println(index, value)
		value = 'x'
	}
	fmt.Println(iter)
}

func TestForRangeEdit(t *testing.T) {
	iter := []int{1, 2, 3, 4, 5, 6}
	for index, value := range iter {
		fmt.Println(index, value)
		iter[index] = 99
	}
	fmt.Println(iter)
}

func TestFor99(t *testing.T) {
	for m := 1; m < 10; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%d x %d = %d ", n, m, m*n)
		}
		fmt.Println()
	}
}

func TestForP(t *testing.T) {

}
