# 函数基础

函数用于对代码块的逻辑封装，提供代码复用的最基本方式

## 定义

Go 语言函数定义格式如下：

```go
func function_name( [parameter list] ) [return_types] {
   函数体
}
```

函数定义解析：

+ func：函数由 func 开始声明
+ function_name：函数名称，函数名和参数列表一起构成了函数签名。
+ parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
+ return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
+ 函数体：函数定义的代码集合。

实例: 求和函数

```go
func sum(x int, y int) int {
    return x + y
}
```

### 类型合并

在声明函数中若存在多个连续形参类型相同可只保留最后一个参数类型名

```go
func sum(x, y int) int {
    return x + y
}
```

### Go 函数可以返回多个值

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}

func main() {
   a, b := swap("Google", "Runoob")
   fmt.Println(a, b)
}
```

### 命名返回值

在函数返回值列表中可指定变量名，变量在调用时会根据类型使用零值进行初始化，在函数
体中可进行赋值，同时在调用 return 时不需要添加返回值，go 语言自动将变量的最终结果
进行返回
在使用命名返回值时，当声明函数中存在若多个连续返回值类型相同可只保留最后一个返回
值类型名

```go
func sum(x, y int) (rest int) {
    rest = x + y
    return
}
```

## 调用

当创建函数时，你定义了函数需要做什么，通过调用该函数来执行指定任务。

调用函数，向函数传递参数，并返回值，例如：

```go
func main() {
    rest := sum(5, 6)
    fmt.Println(rest)
}
```

## 参数

函数如果使用参数，该变量可称为函数的形参。形参就像定义在函数体内的局部变量

### 值传递

值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数

默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200

   fmt.Printf("交换前 a 的值为 : %d\n", a )
   fmt.Printf("交换前 b 的值为 : %d\n", b )

   /* 通过调用函数来交换值 */
   swap(a, b)

   fmt.Printf("交换后 a 的值 : %d\n", a )
   fmt.Printf("交换后 b 的值 : %d\n", b )
}

/* 定义相互交换值的函数 */
func swap(x, y int) int {
   var temp int

   temp = x /* 保存 x 的值 */
   x = y    /* 将 y 值赋给 x */
   y = temp /* 将 temp 值赋给 y*/

   return temp;
}
```

### 引用传递

引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int= 200

   fmt.Printf("交换前，a 的值 : %d\n", a )
   fmt.Printf("交换前，b 的值 : %d\n", b )

   /* 调用 swap() 函数
   * &a 指向 a 指针，a 变量的地址
   * &b 指向 b 指针，b 变量的地址
   */
   swap(&a, &b)

   fmt.Printf("交换后，a 的值 : %d\n", a )
   fmt.Printf("交换后，b 的值 : %d\n", b )
}

func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保存 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}
```

问题: 使用切片的过程分析程序输出结果

```go
//定义一个函数，给切片添加一个元素
func addOne(s []int) {
    s[0] = 4  // 可以改变原切片值
    s = append(s, 1)  // 扩容后分配了新的地址，原切片将不再受影响
    s[0] = 8 
}
var s1 = []int{2}   // 初始化一个切片
addOne(s1)          // 调用函数添加一个切片
fmt.Println(s1)     // 输出一个值 [4]
```

## 匿名函数

## 延迟计算

defer

## 错误处理

panic和recover

## 递归函数

### 内置函数
在 buildin/buildin.go内定义了Go所有支持内置函数：make、len、cap、new、append、copy、close、delete、complex、real、 imag、panic、recover

我们在代码里面可以直接使用: 比如使用len计算字符串的长度
```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println(len("asdfsdf"))
}
```

内置函数和用户定义函数没有本质上的区别, 我们也可以覆盖它
```go
package main

import (
	"fmt"
)

func len(a string) string {
	return "hello"
}

func main() {
	fmt.Println(len("asdfsdf"))
}
```
但是你最好不要这样做，你会被打的