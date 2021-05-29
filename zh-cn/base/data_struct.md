# 数据类型

## 布尔类型

布尔类型用于表示真假，类型名为 bool，只有两个值 true 和 false，占用一个字节宽度，零值为 false

```go
var (
    IsBoy  bool = true
    IsGirl bool = true
)
```

## 数值类型

Go 语言提供了 5 种有符号、5 种无符号、1 种指针、1种单字节、1 种单个unicode字符（unicode
码点），共 13 种整数类型,零值均为 0

+ 整型
|  类型名   | 字节宽度  | 取值范围 |
|  ----  | ----  | --- |
|int     |与平台有关| 32 位系统 4 字节，64 位系统 8 字节 有符号整型|
|uint | 与平台有关|32 位 系统 4 字节，64 位 系统 8 字节 无符号整形|
|int8 |1 字节| 用 8 位表示的有符号整型 取值范围为:[-128, 127]|
|int16 |2 字节| 用 16 位表示的有符号整型  取值范围为：[-32768,32767]|
|int32  | 4 字节 | 用 32 位表示的有符号整型，取值范围为：[-2147483648,2147483647] |
|int64  |8 字节  |用 64 位表示的有符号整型，取值范围为：[-9223372036854775808,9223372036854775807]|
|uint8 |1 字节 | 用 8 位表示的无符号整型，取值范围为：[0,255]|
|uint16 |2 字节 | 用 16 位表示的无符号整型，取值范围为：[0,65535]|
|uint32 |4 字节 | 用 32 位表示的无符号整型，取值范围为：[0,4294967295]|
|uint64 | 8 字节 | 用 64 位表示的无符号整型，取值范围为：[0,18446744073709551615] |
|uintptr |与平台有关|32 位系统 4 字节，64 位系统 8 字节指针值的无符号整型|

+ 浮点型
|  类型名   | 字节宽度  | 取值范围 |
|  ----  | ----  | --- |
| float32| 4字节 | IEEE-754 32位浮点型数|
| float64| 8字节 | IEEE-754 64位浮点型数|
| complex64| 4字节| 32 位实数和虚数|
| complex128| 8字节| 64 位实数和虚数|

+ 别名
|  类型名   | 字节宽度  | 取值范围 |
|  ----  | ----  | --- |
|byte |1 字节| 字节类型，取值范围同 uint8|
|rune |4 字节| Unicode 码点 取值范围同 uint32|

## 字符串类型

Go 语言内置了字符串类型，使用 string 表示

字面量：

+ 可解析字符串：通过双引号(")来创建，不能包含多行，支持特殊字符转义序列
+ 原生字符串：通过反引号(`)来创建，可包含多行，不支持特殊字符转义序列

特殊字符：

+ \\：反斜线
+ \'：单引号
+ \"：双引号
+ \a：响铃
+ \b：退格
+ \f：换页
+ \n：换行
+ \r：回车
+ \t：制表符
+ \v：垂直制表符
+ \ooo：3 个 8 位数字给定的八进制码点的 Unicode 字符（不能超过\377
+ \uhhhh：4 个 16 位数字给定的十六进制码点的 Unicode 字符
+ \Uhhhhhhhh：8 个 32 位数字给定的十六进制码点的 Unicode 字符
+ \xhh：2 个 8 位数字给定的十六进制码点的 Unicode 字符

## 枚举类型

常使用 iota 生成器用于初始化一系列相同规则的常量，批量声明常量的第一个常量使用
iota 进行赋值，此时 iota 被重置为 0，其他常量省略类型和赋值，在每初始化一个常量则
加 1

## 指针类型

每个变量在内存中都有对应存储位置（内存地址），可以通过&运算符获取。指针是用来存储
变量地址的变量

1 声明

指针声明需要指定存储地址中对应数据的类型，并使用*作为类型前缀。指针变量声明
后会被初始化为 nil，表示空指针

2 初始化

+ 使用&运算符+变量初始化：&运算获取变量的存储位置来初始化指针变量
+ 使用 new 函数初始化：new 函数根据数据类型申请内存空间并使用零值填充，并返回申请空间地址

3 操作

可通过*运算符+指针变量名来访问和修改对应存储位置的值

4 指针的指针

用来存储指针变量地址的变量叫做指针的指针

## 自定义类型

在 go 语言中使用 type 声明一种新的类型，语法格式为：

```go
type TypeName Formatter
```

Format 可以时任意内置类型、函数签名、结构体、接口

```go
// 自定义int类型
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}

// Helper handlers

// NotFound replies to the request with an HTTP 404 not found error.
func NotFound(w ResponseWriter, r *Request) { Error(w, "404 page not found", StatusNotFound) }

// NotFoundHandler returns a simple request handler
// that replies to each request with a ``404 page not found'' reply.
func NotFoundHandler() Handler { return HandlerFunc(NotFound) }
```

## 作业

+ 查阅Go源码, 理解AST语法树和Token解析的过程