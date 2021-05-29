# 复合类型

## 数组

数组是具有相同数据类型的数据项组成的一组长度固定的序列，数据项叫做数组的元素，数
组的长度必须是非负整数的常量，长度也是类型的一部分

位置: runtime/slice.go

```go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```

初始化逻辑: makeslice

```go
func makeslice(et *_type, len, cap int) unsafe.Pointer {
    ...
}
```

1 声明
数组声明需要指定组成元素的类型以及存储元素的数量（长度）。在数组声明后，其长度不
可修改，数组的每个元素会根据对应类型的零值对进行初始化

2 字面量
a) 指定数组长度: [length]type{v1, v2, …, vlength}
b) 使用初始化元素数量推到数组长度: […]type{v1, v2, …, vlength}
c) 对指定位置元素进行初始化: [length]type{im:vm, …, sin:in}

3 操作

+ 关系运算==、!=
+ 获取数组长度 使用 len 函数可获取数组的长度
+ 访问&修改 通过对编号对数组元素进行访问和修改，元素的编号从左到右依次为:0, 1, 2, …, n(n为数组长度-1)
+ 切片: array\[start:end\]获取数组的一部分元素做为切片
+ 遍历 可以通过 for+len+访问方式或 for-range 方式对数组中元素进行遍历使用 for-range 遍历数组，range 返回两个元素分别为数组元素索引和值

4 多维数组

数组的元素也可以是数组类型，此时称为多维数组

+ 声明&初始化
+ 访问&修改
+ 遍历

## 切片

切片是长度可变的数组（具有相同数据类型的数据项组成的一组长度可变的序列），切片由
三部分组成：

+ 指针：指向切片第一个元素指向的数组元素的地址
+ 长度：切片元素的数量
+ 容量：切片开始到结束位置元素的数量

### 源码解读

动态增长逻辑: growslice

```go
// growslice handles slice growth during append.
// It is passed the slice element type, the old slice, and the desired new minimum capacity,
// and it returns a new slice with at least that capacity, with the old data
// copied into it.
// The new slice's length is set to the old slice's length,
// NOT to the new requested capacity.
// This is for codegen convenience. The old slice's length is used immediately
// to calculate where to write new values during an append.
// TODO: When the old backend is gone, reconsider this decision.
// The SSA backend might prefer the new length or to return only ptr/cap and save stack space.
func growslice(et *_type, old slice, cap int) slice {
    ...
}
```

1) 声明
切片声明需要指定组成元素的类型，但不需要指定存储元素的数量（长度）。在切片声明后，
会被初始化为 nil，表示暂不存在的切片

2) 初始化
a) 使用字面量初始化:[]type{v1, v2, …, vn}
b) 使用字面量初始化空切片: []type{}
c) 指定长度和容量字面量初始化:[]type{im:vm, in:vn, ilength:vlength}
d) 使用 make 函数初始化
make([]type, len)/make([]type, len, cap)，通过 make 函数创建长度为 len，容量
为 cap 的切片，len 必须小于等于 cap
e) 使用数组切片操作初始化：array[start:end] array[start:end:cap] (end<=cap<=len)

3) 操作
a) 获取切片长度和容量
使用 len 函数可获取切片的长度，使用 cap 函数可获取切片容量

b) 访问和修改
通过对编号对切片元素进行访问和修改，元素的编号从左到右依次为:0, 1, 2, …, n(n
为切片长度-1)

c) 切片: slice[start:end]用于创建一个新的切片，end <= src_cap
新创建切片长度和容量计算：len: end-start, cap: src_cap-start
切片共享底层数组，若某个切片元素发生变化，则数组和其他有共享元素的切片也会发生变化

slice[start:end:cap]可用于限制新切片的容量值, end<=cap<= src_cap
新创建切片长度和容量计算：len: end-start, cap: cap-start

d) 遍历
可以通过 for+len+访问方式或 for-range 方式对切片中元素进行遍历
使用 for-range 遍历切片，range 返回两个元素分别为切片元素索引和值

e) 增加元素
使用 append 对切片增加一个或多个元素并返回修改后切片，当长度在容量范围内时只
增加长度，容量和底层数组不变。当长度超过容量范围则会创建一个新的底层数组并对
容量进行智能运算(元素数量<1024 时，约按原容量 1 倍增加，>1024 时约按原容量 0.25
倍增加)

f) 复制切片到另一个切片
复制元素数量为 src 元素数量和 dest 元素数量的最小值

4) 使用
a) 移除元素

b) 队列
先进先出

c) 堆栈

5) 多维切片
切片的元素也可以是切片类型，此时称为多维切片


先进后出

## 映射

映射是存储一系列无序的 key/value 对，通过 key 来对 value 进行操作（增、删、改、查）。
映射的 key 只能为可使用==运算符的值类型（字符串、数字、布尔、数组），value 可以为
任意类型 s

map.go

```go
// A header for a Go map.
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}
```

```go
// makemap implements Go map creation for make(map[k]v, hint).
// If the compiler has determined that the map or the first bucket
// can be created on the stack, h and/or bucket may be non-nil.
// If h != nil, the map can be created directly in h.
// If h.buckets != nil, bucket pointed to can be used as the first bucket.
func makemap(t *maptype, hint int, h *hmap) *hmap
```

 声明
map 声明需要指定组成元素 key 和 value 的类型，在声明后，会被初始化为 nil，表示
暂不存在的映射

2) 初始化
a) 使用字面量初始化:map[ktype]vtype{k1:v1, k2:v2, …, kn:vn}
b) 使用字面量初始化空映射:map[ktype]vtype{ }
c) 使用 make 函数初始化
make(map[ktype]vtype)，通过 make 函数创建映射

3) 操作
a) 获取元素的数量
使用 len 函数获取映射元素的数量

b) 访问

当访问 key 存在与映射时则返回对应的值，否则返回值类型的零值
c) 判断 key 是否存在
通过 key 访问元素时可接收两个值，第一个值为 value，第二个值为 bool 类型表示元
素是否存在，若存在为 true，否则为 false

d) 修改&增加
使用 key 对映射赋值时当 key 存在则修改 key 对应的 value，若 key 不存在则增加 key
和 value

e) 删除
使用 delete 函数删除映射中已经存在的 key

f) 遍历
可通过 for-range 对映射中个元素进行遍历，range 返回两个元素分别为映射的 key 和
value

4) 使用
统计演讲稿中“我有一个梦想”中各英文字符出现的次数



## 管道

chan.go

```go
type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}
```

```go
func makechan(t *chantype, size int) *hchan
```

