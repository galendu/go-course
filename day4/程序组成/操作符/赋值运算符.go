package main

import "fmt"

func main() {
	var a int
	a = 10 // 修改的是啥? 8Byte 内存空间里面的值
	var b *int
	b = &a // 修改的是啥? 8Byte 内存空间里面的值

	c := 30
	b = &c // b里面存放的就是 c的地址
	fmt.Println(b)

	// 一次赋值多个内存块
	x, y, z := 10, "y", 10.1
	fmt.Println(x, y, z)
}
