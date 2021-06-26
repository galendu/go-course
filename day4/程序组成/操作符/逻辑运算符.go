package main

import "fmt"

func main() {
	a, _ := false, false // /dev/null false  1B<>, _ 就是把值丢掉
	if !a {              // a != true
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	fmt.Println(("x" == "y") && (1 == 1))
	fmt.Println(("x" == "y") || (1 == 1))
	fmt.Println(("x" == "y") && (1 == 1))
	fmt.Println((!("x" == "y")) && (!(1 == 1)))  // () 

	x := 10
	p := &x
	fmt.Println(*p) // 解引用 p(0xffff) --> 10

	// 讲师-老喻 2021/6/26 15:20:03
	// *p
	// (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(p1)) + unsafe.Sizeof(a[0])))  (*T)(unsafe.Pointer)  (int32)(10)  // 单目运算符 uint64
	// unsafe.Pointer(uintptr(unsafe.Pointer(p1)) + unsafe.Sizeof(a[0]))  // T(T(T(T)))  --> (T)()  uintptr  0x1 + ox2 = 0x3
}
