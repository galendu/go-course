package homework

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type Book struct {
	Title  string
	Author string
	Page   uint
	Tag    []string
}

func TestHomeWork1(t *testing.T) {
	b := Book{Tag: []string{"abc", "def", "hjk"}}

	p := (*[]string)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + unsafe.Offsetof(b.Tag)))
	fmt.Println(*p)

	fmt.Printf("%p\n", &b.Tag[0])
	h := (*reflect.SliceHeader)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + unsafe.Offsetof(b.Tag)))
	fmt.Printf("%p\n", h)
	ptr := (*string)(unsafe.Pointer(h.Data))
	fmt.Printf("%p -> %v\n", ptr, *ptr)

	fmt.Printf("%p\n", &b.Tag[1])
	ptr1 := (*string)(unsafe.Pointer(h.Data + 16))
	fmt.Printf("%p -> %v\n", ptr1, *ptr1)

	fmt.Printf("%p\n", &b.Tag[2])
	ptr2 := (*[]byte)(unsafe.Pointer(h.Data + 16))
	fmt.Printf("%p -> %v\n", ptr2, *ptr2)
}

type Student struct {
	Name     string   // 名称
	Number   uint16   // 学号
	Subjects []string // 数学  语文  英语
	Score    []int    //  88   99   77
}

type Class struct {
	Name     string     // 班级名称
	Number   uint8      // 班级编号
	Students []*Student // 班级学员
}

func (c *Class) AvgScore() []int {
	return nil
}
