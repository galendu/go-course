package main

import "fmt"

// 变量就是一个地址, 也就是一个引用，如果我们想要保持这个内存地址应该怎么办?
func main() {

	// 指针的演化, 需要给他一个定义
	a := 10
	fmt.Println(&a) // 0xc0000ac058

	c := &a        //  0xc0000ac058,   指针, *int
	fmt.Println(c) // *T
	fmt.Printf("%T\n", c)
	fmt.Println(&c)

	// 声明
	// *T
	// 没有指向
	var e *int
	var f *float32
	fmt.Println(e, f)

	// 没有指向是, 修改指向的值 是会pannic
	// *e = 10 invalid memory address or nil pointer dereference

	e = new(int) // 内存地址 + 类型约束?
	// new, 开辟一个类型宽带对应的内存空间, 返回一个内存空间的地址
	// make, [makeslice, makemap]  也是内存地址(unsafe.Pointer --> slice struct ---member: pointer---->array)  []int
	fmt.Println(e) // 0xc0000140d0
	// 指针如何赋值

}
