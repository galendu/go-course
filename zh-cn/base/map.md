# Go语言Map

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