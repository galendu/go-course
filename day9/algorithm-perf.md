# 算法的评估

算法（Algorithm）是指用来操作数据、解决程序问题的一组方法。上一小节我们实现了一个插入排序的算法, 当然排序还有很多算法: 
+ 冒泡排序: 两个数比较大小，较大的数下沉，较小的数冒起来
+ 选择排序: 在长度为N的无序数组中，第一次遍历n-1个数，找到最小的数值与第一个元素交换，第二次遍历n-2个数，找到最小的数值与第二个元素交换。。。第n-1次遍历，找到最小的数值与第n-1个元素交换，排序完成
+ 插入排序: 在要排序的一组数中，假定前n-1个数已经排好序，现在将第n个数插到前面的有序数列中，使得这n个数也是排好顺序的。如此反复循环，直到全部排号顺序
+ 快速排序: 快速排序是对冒泡排序的一种改进，也属于交换类的排序算法
+ 其他

## 排序算法

我们挑选一个最简单的冒泡排序来实现 然后对比他们性能:

+ 冒泡排序
+ 选择排序
+ 插入排序
+ 内置排序



### 冒泡排序

基本思想：两个数比较大小，较大的数下沉，较小的数冒起来。具体如下图所示

![](../image/quick-sort.jpeg)


定义我们需要实现的排序算法函数的名称:
```go
func BubbleSort(numbers []int) []int {
    ...
}
```

编写测试用例:
```go
func TestBubbleSort(t *testing.T) {
	should := assert.New(t)

	raw := []int{3, 6, 4, 2, 11, 10, 5}
	target := sort.BubbleSort(raw)

	should.Equal([]int{2, 3, 4, 5, 6, 10, 11}, target)
}
```

接下来大家思考5分钟, 看看能不能自己实现

![](../image/think-kawayi.jpg)

1. 先编写比较的流程
```go
func BubbleSort(numbers []int) []int {
	for i := range numbers {
		for j := 0; j < len(numbers)-1; j++ {
			// 当前值 numbers[i], 后一个值是多少 numbers[j+1]
			fmt.Printf("数据: 当前: %d, 比对: %d\n", numbers[j], numbers[j+1])
		}
		fmt.Printf("第%d趟: \n", i+1)
	}
	return numbers
}
```

看下比对流程是否正确:
```go
数据: 当前: 3, 比对: 6
数据: 当前: 6, 比对: 4
数据: 当前: 4, 比对: 2
数据: 当前: 2, 比对: 11
数据: 当前: 11, 比对: 10
数据: 当前: 10, 比对: 5
第1趟:
数据: 当前: 3, 比对: 6
数据: 当前: 6, 比对: 4
数据: 当前: 4, 比对: 2
数据: 当前: 2, 比对: 11
数据: 当前: 11, 比对: 10
数据: 当前: 10, 比对: 5
第2趟:
数据: 当前: 3, 比对: 6
数据: 当前: 6, 比对: 4
数据: 当前: 4, 比对: 2
数据: 当前: 2, 比对: 11
数据: 当前: 11, 比对: 10
数据: 当前: 10, 比对: 5
第3趟:
数据: 当前: 3, 比对: 6
数据: 当前: 6, 比对: 4
数据: 当前: 4, 比对: 2
数据: 当前: 2, 比对: 11
数据: 当前: 11, 比对: 10
数据: 当前: 10, 比对: 5
第4趟:
数据: 当前: 3, 比对: 6
数据: 当前: 6, 比对: 4
数据: 当前: 4, 比对: 2
数据: 当前: 2, 比对: 11
数据: 当前: 11, 比对: 10
数据: 当前: 10, 比对: 5
第5趟:
数据: 当前: 3, 比对: 6
数据: 当前: 6, 比对: 4
数据: 当前: 4, 比对: 2
数据: 当前: 2, 比对: 11
数据: 当前: 11, 比对: 10
数据: 当前: 10, 比对: 5
第6趟:
数据: 当前: 3, 比对: 6
数据: 当前: 6, 比对: 4
数据: 当前: 4, 比对: 2
数据: 当前: 2, 比对: 11
数据: 当前: 11, 比对: 10
数据: 当前: 10, 比对: 5
第7趟:
```

2. 然后我们补充交换逻辑:

```go
func BubbleSort(numbers []int) []int {
	for i := range numbers {
		for j := 0; j < len(numbers)-1; j++ {
			// 当前值 numbers[i], 后一个值是多少 numbers[j+1]
			fmt.Printf("数据: 当前: %d, 比对: %d\n", numbers[j], numbers[j+1])

			// 比较2个数, 交换顺序, 大数沉底, 小数冒出
			if numbers[j+1] < numbers[j] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
		fmt.Printf("第%d趟: %v\n", i+1, numbers)
	}
	return numbers
}
```

查看比较逻辑
```go
数据: 当前: 3, 比对: 6
数据: 当前: 6, 比对: 4
数据: 当前: 6, 比对: 2
数据: 当前: 6, 比对: 11
数据: 当前: 11, 比对: 10
数据: 当前: 11, 比对: 5
第1趟: [3 4 2 6 10 5 11]
数据: 当前: 3, 比对: 4
数据: 当前: 4, 比对: 2
数据: 当前: 4, 比对: 6
数据: 当前: 6, 比对: 10
数据: 当前: 10, 比对: 5
数据: 当前: 10, 比对: 11
第2趟: [3 2 4 6 5 10 11]
数据: 当前: 3, 比对: 2
数据: 当前: 3, 比对: 4
数据: 当前: 4, 比对: 6
数据: 当前: 6, 比对: 5
数据: 当前: 6, 比对: 10
数据: 当前: 10, 比对: 11
第3趟: [2 3 4 5 6 10 11]
数据: 当前: 2, 比对: 3
数据: 当前: 3, 比对: 4
数据: 当前: 4, 比对: 5
数据: 当前: 5, 比对: 6
数据: 当前: 6, 比对: 10
数据: 当前: 10, 比对: 11
第4趟: [2 3 4 5 6 10 11]
数据: 当前: 2, 比对: 3
数据: 当前: 3, 比对: 4
数据: 当前: 4, 比对: 5
数据: 当前: 5, 比对: 6
数据: 当前: 6, 比对: 10
数据: 当前: 10, 比对: 11
第5趟: [2 3 4 5 6 10 11]
数据: 当前: 2, 比对: 3
数据: 当前: 3, 比对: 4
数据: 当前: 4, 比对: 5
数据: 当前: 5, 比对: 6
数据: 当前: 6, 比对: 10
数据: 当前: 10, 比对: 11
第6趟: [2 3 4 5 6 10 11]
数据: 当前: 2, 比对: 3
数据: 当前: 3, 比对: 4
数据: 当前: 4, 比对: 5
数据: 当前: 5, 比对: 6
数据: 当前: 6, 比对: 10
数据: 当前: 10, 比对: 11
第7趟: [2 3 4 5 6 10 11]
```


### 选择排序

基本思想：在长度为N的无序数组中，第一次遍历n-1个数，找到最小的数值与第一个元素交换，第二次遍历n-2个数，找到最小的数值与第二个元素交换。。。第n-1次遍历，找到最小的数值与第n-1个元素交换，排序完成

![](../image/sort-choice.jpeg)

定义我们需要实现的排序算法函数的名称:
```go
func SelectSort(numbers []int) []int {
	...
}
```

编写测试用例
```go
func TestSelectSort(t *testing.T) {
	should := assert.New(t)

	raw := []int{3, 6, 4, 2, 11, 10, 5}
	target := sort.SelectSort(raw)

	should.Equal([]int{2, 3, 4, 5, 6, 10, 11}, target)
}
```

接下来大家思考5分钟, 看看能不能自己实现

![](../image/think.jpg)


1. 我们先编写数据比较流程

```go
func SelectSort(numbers []int) []int {
	for i := range numbers {
		// 拿到第一个的数, 就是numbers[i], 比如 3
		fmt.Printf("第%d趟: %d\n", i+1, numbers[i])

		// 后后面的数依次比较
		for j := i + 1; j < len(numbers); j++ {
			fmt.Printf("  当前数据: %d, 比对数据: %d\n", numbers[i], numbers[j])
		}
	}
	return numbers
}
```

看下参与比对的数据是否正确:
```
第1趟: 3
  当前数据: 3, 比对数据: 6
  当前数据: 3, 比对数据: 4
  当前数据: 3, 比对数据: 2
  当前数据: 3, 比对数据: 11
  当前数据: 3, 比对数据: 10
  当前数据: 3, 比对数据: 5
第2趟: 6
  当前数据: 6, 比对数据: 4
  当前数据: 6, 比对数据: 2
  当前数据: 6, 比对数据: 11
  当前数据: 6, 比对数据: 10
  当前数据: 6, 比对数据: 5
第3趟: 4
  当前数据: 4, 比对数据: 2
  当前数据: 4, 比对数据: 11
  当前数据: 4, 比对数据: 10
  当前数据: 4, 比对数据: 5
第4趟: 2
  当前数据: 2, 比对数据: 11
  当前数据: 2, 比对数据: 10
  当前数据: 2, 比对数据: 5
第5趟: 11
  当前数据: 11, 比对数据: 10
  当前数据: 11, 比对数据: 5
第6趟: 10
  当前数据: 10, 比对数据: 5
第7趟: 5
```

2. 然后我们补上交换的逻辑
```go
func SelectSort(numbers []int) []int {
	for i := range numbers {
		// 拿到第一个的数, 就是numbers[i], 比如 3
		fmt.Printf("第%d趟: %d\n", i+1, numbers[i])

		// 依次和后面相邻的数比较
		for j := i + 1; j < len(numbers); j++ {
			fmt.Printf("数据 -->  当前数据: %d, 比对数据: %d\n", numbers[i], numbers[j])
			if numbers[i] > numbers[j] {
				// 如果当前数 > 后面的数据, 则交换位置
				numbers[i], numbers[j] = numbers[j], numbers[i]
				fmt.Printf("交换 -->  当前数据: %d, 比对数据: %d\n", numbers[i], numbers[j])
			}
		}

		fmt.Println("结果: ", numbers)
	}

	fmt.Println("最终结果", numbers)
	return numbers
}
```

再次看排序过程
```go
第1趟: 3
数据 -->  当前数据: 3, 比对数据: 6
数据 -->  当前数据: 3, 比对数据: 4
数据 -->  当前数据: 3, 比对数据: 2
交换 -->  当前数据: 2, 比对数据: 3
数据 -->  当前数据: 2, 比对数据: 11
数据 -->  当前数据: 2, 比对数据: 10
数据 -->  当前数据: 2, 比对数据: 5
结果:  [2 6 4 3 11 10 5]
第2趟: 6
数据 -->  当前数据: 6, 比对数据: 4
交换 -->  当前数据: 4, 比对数据: 6
数据 -->  当前数据: 4, 比对数据: 3
交换 -->  当前数据: 3, 比对数据: 4
数据 -->  当前数据: 3, 比对数据: 11
数据 -->  当前数据: 3, 比对数据: 10
数据 -->  当前数据: 3, 比对数据: 5
结果:  [2 3 6 4 11 10 5]
第3趟: 6
数据 -->  当前数据: 6, 比对数据: 4
交换 -->  当前数据: 4, 比对数据: 6
数据 -->  当前数据: 4, 比对数据: 11
数据 -->  当前数据: 4, 比对数据: 10
数据 -->  当前数据: 4, 比对数据: 5
结果:  [2 3 4 6 11 10 5]
第4趟: 6
数据 -->  当前数据: 6, 比对数据: 11
数据 -->  当前数据: 6, 比对数据: 10
数据 -->  当前数据: 6, 比对数据: 5
交换 -->  当前数据: 5, 比对数据: 6
结果:  [2 3 4 5 11 10 6]
第5趟: 11
数据 -->  当前数据: 11, 比对数据: 10
交换 -->  当前数据: 10, 比对数据: 11
数据 -->  当前数据: 10, 比对数据: 6
交换 -->  当前数据: 6, 比对数据: 10
结果:  [2 3 4 5 6 11 10]
第6趟: 11
数据 -->  当前数据: 11, 比对数据: 10
交换 -->  当前数据: 10, 比对数据: 11
结果:  [2 3 4 5 6 10 11]
第7趟: 11
结果:  [2 3 4 5 6 10 11]
最终结果 [2 3 4 5 6 10 11]
```

### 插入排序

为了方便测试我们补充一个NumberStack的初始化方法:
```go
func NewNumberStack(numbers []int) *Stack {
	items := make([]Item, 0, len(numbers))
	for i := range numbers {
		items = append(items, numbers[i])
	}
	return &Stack{
		items: items,
	}
}
```

### 内置排序

go 内置提供了sort函数, 用于对象的排序, 参与排序的对象必须实现比较方法(接口设计的真的吊: 排序的核心逻辑: 比较 交由用户自己定义)
```go
// Sort sorts data.
// It makes one call to data.Len to determine n and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
func Sort(data Interface) {
    ...
}
```

我们自己实现一个IntSlice结构
```go
func NewIntSlice(numbers []int) IntSlice {
	return IntSlice(numbers)
}

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
```

然后提供一个比较函数:
```go
func BuildInSort(numbers []int) []int {
	sort.Sort(IntSlice(numbers))
	return numbers
}
```

## 性能测试

我们通过上面的结果可以知道: 对于同一个问题，使用不同的算法，也许最终得到的结果是一样的，但在过程中消耗的资源和时间却会有很大的区别, 那么我们应该如何去衡量不同算法之间的优劣呢? 


1. 准备用于排序的测试数据, 这里我们随机生成
```go
func generateRandomArray(arrayLen int) []int {
	var a []int
	for i := 0; i < arrayLen; i++ {
		a = append(a, rand.Intn(MAX_RAND_LIMIT))
	}
	return a
}
```

2. 编写冒泡排序的 性能测试用例, 由于我们要分别测试 100 1000 10000 个的排序时间, 我们抽象一个基础函数
```go
func benchmarkBubbleSort(i int, b *testing.B) {
	a := generateRandomArray(i)
	sort.BubbleSort(a)
}
```

3. 编写不通数据量下的性能测试用例
```go
func BenchmarkBubbleSort100(b *testing.B) {
	benchmarkBubbleSort(100, b)
}

func BenchmarkBubbleSort1000(b *testing.B) {
	benchmarkBubbleSort(1000, b)
}

func BenchmarkBubbleSort10000(b *testing.B) {
	benchmarkBubbleSort(10000, b)
}
```

4. 依次类推,  为其他排序算法编写基准测试
```go
func benchmarkSelectSortSort(i int, b *testing.B) {
    ...
}

func benchmarkInsertSortSort(i int, b *testing.B) {
    ...
}

func benchmarkBuildInSort(i int, b *testing.B) {
    ...
}
```

5. 开始我们的性能测试

```
goos: darwin
goarch: amd64
pkg: gitee.com/infraboard/go-course/day9/sort
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkBubbleSort100
BenchmarkBubbleSort100-8      	1000000000	         0.0000288 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort1000
BenchmarkBubbleSort1000-8     	1000000000	         0.001025 ns/op	       0 B/op	       0 allocs/op
BenchmarkBubbleSort10000
BenchmarkBubbleSort10000-8    	1000000000	         0.09919 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectSort100
BenchmarkSelectSort100-8      	1000000000	         0.0000468 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectSort1000
BenchmarkSelectSort1000-8     	1000000000	         0.001055 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectSort10000
BenchmarkSelectSort10000-8    	1000000000	         0.1701 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertSort100
BenchmarkInsertSort100-8      	1000000000	         0.0000355 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertSort1000
BenchmarkInsertSort1000-8     	1000000000	         0.001236 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertSort10000
BenchmarkInsertSort10000-8    	1000000000	         0.1243 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuildInSort100
BenchmarkBuildInSort100-8     	1000000000	         0.0000118 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuildInSort1000
BenchmarkBuildInSort1000-8    	1000000000	         0.0001378 ns/op	       0 B/op	       0 allocs/op
BenchmarkBuildInSort10000
BenchmarkBuildInSort10000-8   	1000000000	         0.001570 ns/op	       0 B/op	       0 allocs/op
```

## 算法评估的维度

主要还是从算法所占用的「时间」和「空间」两个维度去考量。

+ 时间维度：是指执行当前算法所消耗的时间，我们通常用「时间复杂度」来描述。
+ 空间维度：是指执行当前算法需要占用多少内存空间，我们通常用「空间复杂度」来描述

### 时间复杂度



### 空间复杂度




## 参考

+ [golang实现常用排序算法](https://blog.csdn.net/benben_2015/article/details/79231929)
+ [排序算法-快速排序](https://segmentfault.com/a/1190000022288936)
+ [golang 写个快速排序](https://www.jianshu.com/p/7a7ad3af5e25)
+ [时间复杂度与空间复杂度的计算](https://cloud.tencent.com/developer/article/1769988)
+ [算法的时间与空间复杂度](https://zhuanlan.zhihu.com/p/50479555)