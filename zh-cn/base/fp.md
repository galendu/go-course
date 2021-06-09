# 函数式编程

函数式编程，是指忽略（通常是不允许）可变数据（以避免它处可改变的数据引发的边际效应），忽略程序执行状态（不允许隐式的、隐藏的、不可见的状态），通过函数作为入参，函数作为返回值的方式进行计算，通过不断的推进（迭代、递归）这种计算，从而从输入得到输出的编程范式

虽然 functional 并不易于泛型复用，但在具体类型，又或者是通过 interface 抽象后的间接泛型模型中，它是改善程序结构、外观、内涵、质量的最佳手段。
所以你会看到，在成熟的类库中，无论是标准库还是第三方库，functional 模式被广泛地采用

例子: 迭代处理

```go
package main

import "fmt"

func main() {
	var list = []int{1, 2, 3, 4}
	// we are passing the array and a function as arguments to mapForEach method.
	var out = mapForEach(list, func(it int) int {
		return pow(it, 3)
	})
	fmt.Println(out) // [1 8 27 64]
}

func pow(x, n int) int {
	ret := 1 // 结果初始为0次方的值，整数0次方为1。如果是矩阵，则为单元矩阵。
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}

// The higher-order-function takes an array and a function as arguments
func mapForEach(arr []int, fn func(it int) int) []int {
	var newArray = []int{}
	for _, it := range arr {
		// We are executing the method passed
		newArray = append(newArray, fn(it))
	}
	return newArray
}
```

## 闭包

## 递归

## 算子

## 可变参数：Functional Options

