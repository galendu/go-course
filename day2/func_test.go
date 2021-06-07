package day2

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func max(a, b int, args ...int) int {
	// 输出args中保存的参数
	for index, value := range args {
		fmt.Printf("%s%d%s %d\n", "args[", index, "]:", value)
	}

	// 取出a、b中较大者
	max_value := a
	if a > b {
		max_value = a
	}

	// 取出所有参数中最大值
	for _, value := range args {
		if max_value < value {
			max_value = value
		}
	}
	return max_value
}

func TestFuncArgs(t *testing.T) {
	fmt.Println(max(1, 2, 3, 4, 5, 9, 10))
}

// n*(n-1)*...*3*2*1, 5*4*3*2*1
func fact(n int) int {
	// 判断退出点
	if n == 1 {
		return 1
	}

	// 递归表达式
	return n * fact(n-1)
}

func fact2(n int) int {
	current := n
	for n > 1 {
		fmt.Printf("%d * (%d-1)\n", current, n)
		current = current * (n - 1)
		n--
	}

	return current
}

func TestFact(t *testing.T) {
	fmt.Println(fact(5))
	fmt.Println(fact2(5))
}

// f(n)=f(n-1)+f(n-2)且f(2)=f(1)=1, 1 2 3 5	8 13
func fib(n int) int {
	// 退出点判断
	if n == 1 || n == 2 {
		return 1
	}
	// 递归表达式
	return fib(n-1) + fib(n-2)
}

func TestFib(t *testing.T) {
	for i := 2; i < 10; i++ {
		fmt.Printf("%d\t", fib(i))
	}
	fmt.Println()
}

func ParseInt(str string) (result int64, err error) {
	result = 10
	return 20, nil
}

func TestRetrun(t *testing.T) {
	fmt.Println(ParseInt(""))
}

func TestReturnHull(t *testing.T) {
	i, _ := strconv.ParseInt("10", 10, 64)
	fmt.Println(i)
}

func TestAnonymousFunc(t *testing.T) {
	a := func(x, y int) int {
		return x + y
	}

	fmt.Println(a(1, 2))
}

func TestAnonymousFunc2(t *testing.T) {
	a := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(a)
}

type addFunc func(x, y int) int

func TestFuncType(t *testing.T) {
	addFunc := func(x, y int) int {
		return x + y
	}
	fmt.Println(reflect.TypeOf(addFunc))
	fmt.Println(addFunc(10, 20))
}
