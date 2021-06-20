# 函数式编程

函数式编程，是指忽略（通常是不允许）可变数据（以避免它处可改变的数据引发的边际效应），忽略程序执行状态（不允许隐式的、隐藏的、不可见的状态），通过函数作为入参，函数作为返回值的方式进行计算，通过不断的推进（迭代、递归）这种计算，从而从输入得到输出的编程范式

虽然 functional 并不易于泛型复用，但在具体类型，又或者是通过 interface 抽象后的间接泛型模型中，它是改善程序结构、外观、内涵、质量的最佳手段。
所以你会看到，在成熟的类库中，无论是标准库还是第三方库，functional 模式被广泛地采用

下面介绍几种关于使用函数式编程的编程模式


## Map-Reduce

这是解耦数据结构和算法最常见的方式

### Map

模式:
```
item1 --map func--> new1
item2 --map func--> new2
item3 --map func--> new3
...
```

比如 我们写一个Map函数来将所有的字符串转换成大写
```go
func TestMap(t *testing.T) {
	list := []string{"abc", "def", "fqp"}
	out := MapStrToUpper(list, func(item string) string {
		return strings.ToUpper(item)
	})
	fmt.Println(out)
}

func MapStrToUpper(data []string, fn func(string) string) []string {
	newData := make([]string, 0, len(data))
	for _, v := range data {
		newData = append(newData, fn(v))
	}

	return newData
}
```


### Reduce

模式:
```
item1 --|
item2 --|--reduce func--> new item
item3 --|
```

比如写一个Reduce函数用于求和
```go
func TestReduce(t *testing.T) {
	list := []string{"abc", "def", "fqp", "abc"}
	// 统计字符数量
	out1 := ReduceSum(list, func(s string) int {
		return len(s)
	})
	fmt.Println(out1)
	// 出现过ab的字符串数量
	out2 := ReduceSum(list, func(s string) int {
		if strings.Contains(s, "ab") {
			return 1
		}
		return 0
	})
	fmt.Println(out2)
}

func ReduceSum(data []string, fn func(string) int) int {
	sum := 0
	for _, v := range data {
		sum += fn(v)
	}
	return sum
}
```

reduce 还有一种比较常用的模式：过滤器
```go
func TestFilter(t *testing.T) {
	list := []string{"abc", "def", "fqp", "abc"}
	out := ReduceFilter(list, func(s string) bool {
		return strings.Contains(s, "f")
	})
	fmt.Println(out)
}

func ReduceFilter(data []string, fn func(string) bool) []string {
	newData := []string{}
	for _, v := range data {
		if fn(v) {
			newData = append(newData, v)
		}
	}
	return newData
}
```

### 应用

比如我们有这样一个数据集合:
```go
type Class struct {
	Name     string     // 班级名称
	Number   uint8      // 班级编号
	Students []*Student // 班级学员
}

type Student struct {
	Name     string   // 名称
	Number   uint16   // 学号
	Subjects []string // 数学  语文  英语
	Score    []int    //  88   99   77
}
```


## 修饰器

## Pipeline

## 算子

## Functional Options


