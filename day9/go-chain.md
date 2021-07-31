# 数据结构之链表

场景: 我们有100个对象需要存储, 最简单的方式是直接使用Array
```go
var store = [100]*T{}
// 存储
store[0] = OBJ
// 变量
for i := range store
```
我不知道需要存多少个, 可能100 也可能200 或者更多喃?

使用切片:
```go
var store = []*T{}
// 追加
store = append(store, *T)
// 访问
for i := range store
```
如果有人要插队, 比如我要在 index2 和 index3 中间插入一个元素?

使用切片来处理插入
```go
var store = []*T{OJB1, OJB2, OBJ3}

// 此时我需要把OJB4 插入到 OJB2 和 OBJ3直接
// 1. 获取之前的和或者之后的
newStore = store = []*T{}

// 拷贝第一部分
newStore = append(newStore, store[:2])

// 2. 添加插入的元素
newStore = append(newStore, OBJ4)
// 3. 拷贝之后的
newStore = append(newStore, store[2:])

// 当然你可以简洁的写成这样
newStore = append(store[:2]..., OBJ4, store[2:]...)
```

我们可以看到 我们插入的位置，会导致后面的数据都往后移动一位, 假设我们有1亿条数据, 我要插入到第一个, 可以想象这个效率会有多低(如果你的场景是从一头那数据，那么栈更高效)

同理删除也一样, 那我们如何改进?

思考:

由于slice底层使用数组存储元素, 受限于数组的内存结构(连续内存空间):
```
A + SIZE --> B + SIZE --> C + SIZE ---> D ....
``` 

那如果我们插入一个元素理论上应该是啥样的最直接
```
A      --->  B       -->  C      --> D

```



## 链表

链表是一种数据结构，和数组不同，链表并不需要一块连续的内存空间，它通过「指针」将一组零散的内存块串联起来使用





## 比较list和slice的插入速度




## 链表怎么好， 我都使用链表

比较slice 与 list 遍历创建和添加元素速度


比较list和slice的遍历速度




## 总结

对于很多数据来讲：
+ 频繁的插入和删除用list
+ 频繁的遍历查询选slice