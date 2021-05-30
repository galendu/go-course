# 基本语法

如下是一个程序的基本组成结构:
```go
// 当前程序的包名, main包表示入口包, 是编译构建的入口
package main

// 导入其他包
import "fmt"

// 常量定义
const PI = 3.1415

// 全局变量声明和赋值
var name = "fly"

// 一般类型声明
type newType int

// 结构体声明
type student struct{}

// 接口声明
type reader interface{}

// 程序入口
func main() {
        fmt.Println("hello world, this is my first golang program!")
}
```

在这个结构里面必须要符合Go程序的语法, 编写出来的程序才是合法的，才能被编译器识别, 编译器识别代码的基础单位是Lexical Token(词法标记)，比如 如下一段代码:
```go
func main() {
        fmt.Println("hello world, this is my first golang program!")
}
```
他包含的Lexical Token(词法标记)，有如下12个:
```
func     // 关键字 func
main     // 标识符 函数名称
(        // LPAREN 左小括号
{        // LBRACE 左花括号
fmt      // 标识符  包名称  
.        // PERIOD 调用符
Println  // 标识符 函数名称
(        // LPAREN 左小括号
"hello world, this is my first golang program!" // 标识符 字符串常量
)        // RPAREN 右小括号
}        // RBRACE 右花括号
)        // LPAREN 右小括号
```

go 语言支持的所有词法标记如下:
```go
var tokens = [...]string{
        // 特殊词法
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	COMMENT: "COMMENT",

        // 标识符
	IDENT:  "IDENT",

        // 基本类型
	INT:    "INT",
	FLOAT:  "FLOAT",
	IMAG:   "IMAG",
	CHAR:   "CHAR",
	STRING: "STRING",

        // 操作符
	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	REM: "%",

	AND:     "&",
	OR:      "|",
	XOR:     "^",
	SHL:     "<<",
	SHR:     ">>",
	AND_NOT: "&^",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=",
	REM_ASSIGN: "%=",

	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",

	LAND:  "&&",
	LOR:   "||",
	ARROW: "<-",
	INC:   "++",
	DEC:   "--",

	EQL:    "==",
	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	NOT:    "!",

	NEQ:      "!=",
	LEQ:      "<=",
	GEQ:      ">=",
	DEFINE:   ":=",
	ELLIPSIS: "...",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",

        // 25个关键字
	BREAK:    "break",
	CASE:     "case",
	CHAN:     "chan",
	CONST:    "const",
	CONTINUE: "continue",

	DEFAULT:     "default",
	DEFER:       "defer",
	ELSE:        "else",
	FALLTHROUGH: "fallthrough",
	FOR:         "for",

	FUNC:   "func",
	GO:     "go",
	GOTO:   "goto",
	IF:     "if",
	IMPORT: "import",

	INTERFACE: "interface",
	MAP:       "map",
	PACKAGE:   "package",
	RANGE:     "range",
	RETURN:    "return",

	SELECT: "select",
	STRUCT: "struct",
	SWITCH: "switch",
	TYPE:   "type",
	VAR:    "var",
}
```

下面我们从上到下依次介绍这些语法

## 特殊词法

+ ILLEGAL： 标识非法的词法
比如如下一段代码:
```go
func main() {
   中文
}
```
这个`中文`就是一个非法的词法, 编译的时候回直接报
```sh
src\day1\hello.go:8:2: undefined: 中文
```

+ EOF： 用于标识 流(io stream)结束
比如parser/parser.go会重复解析声明到文件的最后:
```go
for p.tok != token.EOF {
    decls = append(decls, p.parseDecl(declStart))
}
```

+ COMMENT：注释
Go 支持两种注释方式，行注释和块注释：
行注释：以//开头，例如： //我是行注释
```go
// 这是一个行注释
```
块注释：以/*开头，以*/结尾，例如：/*我是块注释*/
```go
/*
这是一个块注释
*/
```

## 标识符

标识符用来命名变量、类型等程序实体，标识符一般有如下几大类:
+ 命名标识: 变量、常量、函数、类型、接口、包名
+ 内部标识: 内置函数


### 命名标识

一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字

Go 语言标识符的命名规则如下：
+ 只能由非空字母(Unicode)、数字、下划线(_)组成
+ 只能以字母或下划线开头
+ 不能 Go 语言关键字
+ 避免使用 Go 语言预定义标识符
+ 建议使用驼峰式
+ 标识符区分大小写

下面这些就是一些合法的标识符
```sh
username   xxx   M   user_name   user1
_temp   temp_   heelo1  MMXXX  中文
```

比如这段代码是合法的
```go
package main

import (
	"fmt"
)

func main() {
	中文 := "你好，中文"
	fmt.Println(中文)
}
```

而下面这些就是一个非法的标识符
```
1user  // 数字打头
for    // 关键字不能作为标识符
m*m   // 运算符是不允许的
中 午 // 有空格
```

下面这段代码就会报错
```go
package main

import (
	"fmt"
)

func main() {
	m*2 := "stirng"
	fmt.Println(m*2)
}
```

### 内置函数
在 buildin/buildin.go内建议了Go所有的内置函数：make、len、cap、new、append、copy、close、delete、complex、real、 imag、panic、recover

我们在代码里面可以直接使用: 使用len计算字符串的长度
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

## 关键字

关键字用于特定的语法结构，Go 语言定义 25 关键字：

+ 声明：import、package
+ 实体声明和定义：chan、const、func、interface、map、struct、type、var
+ 流程控制：break、case、continue、default、defer、else、fallthrough、for、go、goto、
if、range、return、select、switch

## 操作符

+ 算术运算符：+、-、*、/、%、++、--
+ 关系运算符：>、>=、<、<=、==、!=
+ 逻辑运算符：&&、||、!
+ 位运算符：&、|、^、<<、>>、&^
+ 赋值运算符：=、+=、-=、*=、/=、%=、&=、|=、^=、<<=、>>=
+ 其他运算符：&(单目)、*(单目)、.(点)、-(单目)、…、<-

## 变量

变量是指对一块存储空间定义名称，通过名称对存储空间的内容进行访问或修改，使用 var
进行变量声明，常用的语法为:

+ var 变量名 变量类型 = 值 定义变量并进行初始化，例如：var name string = "silence"
+ var 变量名 变量类型 定义变量使用零值进行初始化，例如：var age int
+ var 变量名 = 值 定义变量，变量类型通过值类型进行推导 例如： var isBoy = true
+ var 变量名 1, 变量名 2 , …, 变量名 n 变量类型 定义多个相同类型的变量并使用零值进行初始化 例如：var prefix, suffix string
+ var 变量名 1, 变量名 2 , …, 变量名 n 变量类型 = 值 1, 值 2, …, 值 n 定义多个相同类型的变量并使用对应的值进行初始化，例如：var prev, next int = 3, 4
+ var 变量名 1, 变量名 2 , …, 变量名 n = 值 1, 值 2, …, 值 n定义多个变量并使用对应的值进行初始化，变量的类型使用值类型进行推导，类型可不相同，例如：var name, age = "silence", 30
+ 批量定义: 定义多个变量并进行初始化，批量复制中变量类型可省略, 初始化表达式可以使用字面量、任何表达式、函数
    var (
    变量名 1 变量类型 1 = 值 1
    变量名 2 变量类型 2 = 值 2
    )

```go
// 例子
var (
name string = "silence"
age int = 30
)
```

## 常量

常量用于定义不可被修改的的值，需要在编译过程中进行计算，只能为基础的数据类型布尔、
数值、字符串，使用 const 进行常量声明，常用语法:

+ const 常量名 类型 = 值 定义常量并进行初始化，例如：const pi float64 = 3.1415926
+ const 常量名 = 值 定义常量，类型通过值类型进行推导，例如：const e = 2.7182818
+ 批量定义: const (常量名 1 类型 1 = 值 1 常量名 2 类型 2 = 值 2)

```go
// 例如：
const (
name string = "silence"
age int = 30
)
const (
name string = "silence"
desc
)
```

定义多个变量并进行初始化，批量复制中变量类型可省略，并且除了第一个常量值外其他常量可同时省略类型和值，表示使用前一个常量的初始化表达式

常量之间的运算，类型转换，以及对常量调用函数 len、cap、real、imag、complex、unsafe.Sizeof 得到的结果依然为常量

## 字面量

字面量是值的表示方法，常用与对变量/常量进行初始化，主要分为：

+ 标识基础数据类型值的字面量，例如：0, 1.1, true, 3 + 4i, 'a', "我爱中国"
+ 构造自定义的复合数据类型的类型字面量，例如：type Interval int
+ 用于表示符合数据类型值的复合字面量，用来构造 array、slice、map、struct 的值，
例如：{1, 2, 3}

## 作用域

作用域指变量可以使用范围。go 语言使用大括号显示的标识作用域范围，大括号内包含一
连串的语句，叫做语句块。语句块可以嵌套，语句块内定义的变量不能在语句块外使用
常见隐式语句块：

+ 全语句块
+ 包语句块
+ 文件语句块
+ if、switch、for、select、case 语句块
作用域内定义变量只能被声明一次且变量必须使用，否则编译错误。在不同作用域可定义相
同的变量，此时局部将覆盖全局

## 注释


