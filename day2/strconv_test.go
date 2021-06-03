package day2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTypeOf1(t *testing.T) {
	a := 10
	b := 0.1314
	c := "hello"

	fmt.Printf("a type: %v\n", reflect.TypeOf(a))
	fmt.Printf("b type: %v\n", reflect.TypeOf(b))
	fmt.Printf("c type: %v\n", reflect.TypeOf(c))
}

func TestF2Int(t *testing.T) {
	a := 3.14
	fmt.Println(int(a))
}

func TestTypeOf(t *testing.T) {
	type Age int
	var a Age = 10
	var b int = 20
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}

func TestCustom(t *testing.T) {
	// Age 底层数据结构为 int
	type Age int
	// a 类型是Age 底层为 int 10
	var a Age = 10

	// 将a转化成int类型,
	// 由于a是Age, 转化成int后, 他们不是同一种类型，不能再次赋值回去: a = int(a) 是不行的
	b := int(a)
	// 现在b是int类型
	fmt.Println(reflect.TypeOf(b))

	// 反过来我们也可以将int类型转换为Age类型
	c := Age(10)
	// 现在c就是Age类型，而不是int类型了
	fmt.Println(reflect.TypeOf(c))
}
