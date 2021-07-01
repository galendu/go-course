# 面向对象

![](../../image/oop.png)

面向对象编程——Object Oriented Programming，简称OOP，是一种程序设计思想。OOP把对象作为程序的基本单元，一个对象包含了数据和操作数据的函数

## 面向过程与面向对象

面向过程的程序设计把计算机程序视为一系列的命令集合，即一组函数的顺序执行。为了简化程序设计，面向过程把函数继续切分为子函数，即把大块函数通过切割成小块函数来降低系统的复杂度。

![](../../image/opp-flow.jpg)

而面向对象的程序设计把计算机程序视为一组对象的集合，而每个对象都可以接收其他对象发过来的消息，并处理这些消息，计算机程序的执行就是一系列消息在各个对象之间传递

![](../../image/message-passing-in-oop.png)

我们以一个例子来说明面向过程和面向对象在程序流程上的不同之处: 求年级学科平均分
```go
type Student struct {
	Name     string   // 名称
	Number   uint16   // 学号  2 ^ 16
	Subjects []string // 数学  语文  英语
	Score    []int    //  88   99   77
}
```

1.面向过程

我们处理逻辑核心部分是函数, 比如会这样写:
```go
func GradeAvg([]*Student) []int {}
```

2.面向对象

如果采用面向对象的程序设计思想，我们首选思考的不是程序的执行流程, 而是年级这种数据类型应该被视为一个对象,
这个对象拥有Students和一些其他属性（Property）, 如果要求年级的平均分, 首先是创建一个年级对应的对象,比如:

```go
// Class 保存的是班级的信息
type Grade struct {
	Number   uint8      // 年级编号
	Subjects []string   // 数学  语文  英语
	Students []*Student // 班级学员, []int --> [10, 20, 30]  []*int ---> [0xaabb, 0xccc, oxddd]
}
```

然后，给对象发一个GradeAvg消息，让对象自己把自己把年级的学科平均值告诉你, 比如:

```go
g := &Grade{}
g.GradeAvg()
```

给对象发消息实际上就是调用对象对应的关联函数，我们称之为对象的方法（Method）。比如:
```go
func (g *Grade) GradeAvg() []int {}
```

面向对象的程序写出来就像这样:
```go
g := &Grade{}
g.GradeAvg()
```

## 类和实例

面向对象的设计思想是从自然界中来的，因为在自然界中, 每一个实体都是对象(Object/Instance), 而这种实体的抽象类别就是类(Class), 比如车就是一个类, 而从你面前路过的福特汽车就是一个实例(Object)
![](../../image/class-object.png)

## Go语言如何面向对象

其实 GO 并不是一个纯面向对象编程语言。它没有提供类（class）这个关键字，只提供了结构体（struct）类型

java 或者 C# 里面，结构体（struct）是不能有成员函数的。然而，Go 语言中的结构体（struct）可以有” 成员函数”。方法可以被添加到结构体中，类似于一个类的实现

## 封装


## 继承


## 多态

