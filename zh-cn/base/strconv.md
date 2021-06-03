# 数据类型转换

Go语言是强类型, 如果需要对数据进行类型转换，需要手动进行, 转换数据类型的简单方法是 直接通过`类型()`方式:
```go
valueOfTypeB = typeB(valueOfTypeA)
```

例如:
```go
a := 3.14
b := int(a)
```

我们可以通过： 标准库 reflect包中的TypeOf方法查看一个变量的类型, 比如:
```go
a := 10
b := 0.1314
c := "hello"

fmt.Printf("a type: %v\n", reflect.TypeOf(a))
fmt.Printf("b type: %v\n", reflect.TypeOf(b))
fmt.Printf("c type: %v\n", reflect.TypeOf(c))
```

## 自定义类型转换

Go允许我们通过type定义自己的类型,自己定义的类型和该类型不是同一类型了，比如:
```go
type Age int
var a Age = 10
var b int = 20
fmt.Println(reflect.TypeOf(a))  // day2.Age
fmt.Println(reflect.TypeOf(b))  // int
```
此时我们定义的Age类型已经不再是int类型了, 只是该类型底层的值为int

```go
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
```

不是所有数据类型都能转换的，例如字母格式的string类型"abcd"转换为int肯定会失败
低精度转换为高精度时是安全的，高精度的值转换为低精度时会丢失精度。例如int32转换为int16，float32转换为int
这种简单的转换方式不能对int(float)和string进行互转，要跨大类型转换，可以使用strconv包提供的函数


