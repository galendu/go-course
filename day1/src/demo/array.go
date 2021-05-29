package main

import "fmt"

func main() {
	var arr1 [2]string // [2]type
	arr1 = [2]string{"1", "2"}
	fmt.Println(arr1)

	var slice1 []string // []type     array1 := [4]string{}
	slice1 = []string{"a", "b", "c"}
	slice1 = append(slice1, "d")
	fmt.Println(slice1)

	var slice2 []string
	slice2 = make([]string, 3, 10) // ["", "", ""]  [10]string
	slice2 = append(slice2, "f")
	fmt.Println(slice2)

	var slice3 []string
	slice3 = make([]string, 0, 2)          // ["", "", ""]  [10]string
	slice3 = append(slice3, "a", "b", "c") // [4]string
	fmt.Println(slice3)

	var slice4 []int
	slice4 = make([]int, 4, 8)
	fmt.Println(slice4, len(slice4))

	// 切片的 切片
	var slice5 []int
	slice5 = []int{1, 2, 3, 4} // [0<1>, 1<2>, 2<3>, 3<4>]
	slice6 := slice5[:]        // [)
	fmt.Println(slice6)

	// 切片复制
	fmt.Println(slice5[1])
	var slice7 []int
	rest := copy(slice7, slice5)
	fmt.Println(rest)
	fmt.Println(slice7)
}
