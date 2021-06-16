package day3

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSafePointer1(t *testing.T) {
	str := "pointer_test"
	a := &str
	fmt.Println(a)
}

func TestPonterT(t *testing.T) {
	var x int64 = 20
	a := &x
	fmt.Println(a)
	// 没有办法指针a的值进行如下转换: *int64 --> *uint64
}

func TestPonterT1(t *testing.T) {
	var x int64 = 20
	a := &x
	fmt.Println(a)

	var y uint64 = 20
	b := &y
	fmt.Println(b)

	// 我们不能进行 a = b
}

type Man struct {
	Name string
	Age  int64
}

func TestUnsafePonter1(t *testing.T) {
	m := Man{Name: "John", Age: 20}
	fmt.Println(unsafe.Sizeof(m.Name), unsafe.Sizeof(m.Age), unsafe.Sizeof(m)) // 16 8 24

	fmt.Println(unsafe.Offsetof(m.Name)) // 0
	fmt.Println(unsafe.Offsetof(m.Age))  // 16
}

func TestUnsafePointer2(t *testing.T) {
	a := [3]int64{1, 2, 3}
	fmt.Printf("%p\n", &a)

	s1 := unsafe.Sizeof(a[0])
	fmt.Printf("%d\n", s1)

	p1 := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + s1))
	fmt.Println(*p1)
}

func TestUnsafePointer3(t *testing.T) {
	type T struct{ a int }
	var t1 T
	fmt.Printf("%p\n", &t1)                          // 0xc0000a0200
	println(&t1)                                     // 0xc0000a0200
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t1))) // c0000a0200
}

type T struct {
	x bool
	y [3]int16
}

const (
	N = unsafe.Offsetof(T{}.y)
	M = unsafe.Sizeof(T{}.y[0])
)

func TestUnsafePointer4(t *testing.T) {
	t1 := T{y: [3]int16{123, 456, 789}}
	p := unsafe.Pointer(&t1)
	// "uintptr(p) + N + M + M"为t.y[2]的内存地址。
	ty2 := (*int16)(unsafe.Pointer(uintptr(p) + N + M + M))
	fmt.Println(*ty2) // 789
}
