# Go语言结构体

![](../../image/struct_title.png)

我们前面介绍的数组 只能保存同一种类型的数据, 当我们需要记录多种不同类型的数据时，我们该怎么办?

结构体就是用于解决这个问题的, 结构体是由一系列具有相同类型或不同类型的数据构成的数据集合, 方便容量我们的任意类型的数据

## 结构体的定义

结构体定义使用 struct 标识，需要指定其包含的属性（名和类型），在定义结构体时可以为
结构体指定结构体名（命名结构体），用于后续声明结构体变量使用

```go
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
```

例如 我们用于描述一个人的特征:

如果我们单独使用变量描述应该如何描述:

```go
var (
    name          string
    age           int
    gender        string
    weight        uint
    favoriteColor []string
)
```



```go
type Person struct {
    Name          string
    Age           int
    Gender        string
    Weight        uint
    FavoriteColor []string
}
```

## 声明与初始化

声明结构体变量只需要定义变量类型为结构体名，变量中的每个属性被初始化为对应类型的
零值。

遵循所有类型声明语法: var struct_name struct_type

使用结构体创建的变量叫做对应结构体的实例或者对象

1.只声明不初始化

比如下面我们初始化一个person的实例

```go
// 只声明
var person Person
```

我们可以看到声明后的结构体的所有属性都是初始值

```go
var person Person
fmt.Printf("%+v\n", person)
// {Name: Age:0 Gender: Weight:0 FavoriteColor:[]}
```

如果我要声明并初始化喃？

2.声明并初始化

```go
var person Person = Person{
    Name:          "andy",
    Age:           66,
    Gender:        "male",
    Weight:        120,
    FavoriteColor: []string{"red", "blue"},
}
fmt.Printf("%+v\n", person)
// {Name:andy Age:66 Gender:male Weight:120 FavoriteColor:[red blue]}
```

注意，上面最后一个逗号","不能省略，Go会报错，这个逗号有助于我们去扩展这个结构

## 属性的访问和修改

通过结构体对象名.属性名的方式来访问和修改对象的属性值

可以通过结构体指针对象的点操作直接对对象的属性值进行访问和修改

```go
结构体.成员名
```

```go
```

## 结构体指针

1.声明

和其他基础数据类型一样，我们也可声明结构体指针变量，此时变量被初始化为 nil

```go
var person *Person
fmt.Println(person)
// <nil>
```

2.声明并初始化

```go
var person *Person = &Person{
    Name:          "andy",
    Age:           66,
    Gender:        "male",
    Weight:        120,
    FavoriteColor: []string{"red", "blue"},
}
fmt.Printf("%p", person)
```

3.new函数创建指针对象

Go 语言中常定义 N(n)ew+结构体名命名的函数用于创建对应的结构体值对象或指针对象

```go
person := new(Person)
fmt.Printf("%p", person)
```

## 结构体方法

属于数据结构的函数，可以为数据结构定义属于自己的函数



## 结构体嵌套

复杂从此开始

1.匿名嵌套

在定义变量时将类型指定为结构体的结构，此时叫匿名结构体。匿名结构体常用于初始化一
次结构体变量的场景，例如项目配置

匿名结构体可以组合不同类型的数据，使得处理数据变得更为灵活。尤其是在一些需要将多个变量、类型数据组合应用的场景，匿名结构体是一个不错的选择

```go
结构体.成员名
```

```go
package main
 
import (
	"encoding/json"
	"fmt"
)
//定义手机屏幕
type Screen01 struct {
	Size       float64 //屏幕尺寸
	ResX, ResY int //屏幕分辨率 水平 垂直
}
//定义电池容量
type Battery struct {
	Capacity string
}
 
//返回json数据
func getJsonData() []byte {
	//tempData 接收匿名结构体（匿名结构体使得数据的结构更加灵活）
	tempData := struct {
		Screen01
		Battery
		HashTouchId bool  // 是否有指纹识别
	}{
		Screen01:    Screen01{Size: 12, ResX: 36, ResY: 36},
		Battery:     Battery{"6000毫安"},
		HashTouchId: true,
	}
	jsonData, _ := json.Marshal(tempData)  //将数据转换为json
	return jsonData
}
```

2命名嵌套

结构体命名嵌入是指结构体中的属性对应的类型也是结构体

适用于复合数据结构<嵌入匿名>

1. 定义

```go
type Book struct {
    Author  struct{
        Name string
        Aage int
    }
    Titile struct{
        Main string 
        Sub  string
    }
}
```

2.声明和初始化

```go
b := &Book{
    Author: struct {
        Name string
        Aage int
    }{
        Name: "xxxx",
        Aage: 11,
    },
    Titile: struct {
        Main string
        Sub  string
    }{
        Main: "xxx",
        Sub:  "yyy",
    },
}

// 不会有人愿意那样用的
b := new(Book)
b.Author.Aage = 11
b.Author.Name = "xxx"
```

2.嵌入命名

```go
type Author struct {
    Name string
    Aage int
}

type Titile struct {
    Main string
    Sub  string    
}

type Book struct {
    Author Author
    Titile Titile
}
```

声明

```go
b := &Book{
    Author: Author{
        Name: "xxx",
        Aage: 11,
    },
    Titile: Titile{
        Main: "t1",
        Sub:  "t2",
    },
}
```

3.属性的访问和修改

```go
b.Author.Name = "xxx"
```

## 结构体的组合: 嵌入


1.匿名嵌入

结构体匿名嵌入是指将已定义的结构体名直接声明在新的结构体中，从而实现对以后已有类
型的扩展和修改

1. 定义

2. 声明&初始化

3. 属性访问和修改

2.指针类型嵌入

1.定义

2.声明&初始化&操作

## 可见性

结构体首字母大写则包外可见(公开的)，否者仅包内可访问(内部的)
结构体属性名首字母大写包外可见(公开的)，否者仅包内可访问(内部的)
组合：

+ 结构体名首字母大写，属性名大写：结构体可在包外使用，且访问其大写的属性名
+ 结构体名首字母大写，属性名小写：结构体可在包外使用，且不能访问其小写的属性名
+ 结构体名首字母小写，属性名大写：结构体只能在包内使用，属性访问在结构体嵌入时 由被嵌入结构体(外层)决定，被嵌入结构体名首字母大写时属性名包外可见，否者只能
在包内使用
+ 结构体名首字母小写，属性名小写：结构体只能在包内使用

## 拷贝

有时候 为了不让对象直接相互干扰, 我们需要深度赋值对象

1. 浅拷贝



2. 深拷贝
