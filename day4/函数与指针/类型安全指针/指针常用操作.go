package main

import "fmt"

func main() {
	// 取地址的值  a(a变量的地址 0xffff) ----> *int 也是一个地址 --->   int(8B)值 0
	a := new(int)
	fmt.Println(&a, a, *a) // 0xc000006028 0xc000014098 0

	// & 获取变量的内存地址
	// * 读取地址的值
}
