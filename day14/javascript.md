# Javascript基础

需要为HTML页面上添加一些动态效果, Brendan Eich这哥们在两周之内设计出了JavaScript语言

几个公司联合ECMA（European Computer Manufacturers Association）组织定制了JavaScript语言的标准，被称为ECMAScript标准

## JavaScript 运行时

+ 浏览器
+ NodeJS

```sh
$ node -v
v14.17.1
```

## 数据类型

+ Number
+ 字符串
+ 布尔值
+ 数组
+ 对象

### 数组

JavaScript的Array可以包含任意数据类型，并通过索引来访问每个元素

```js
var arr1 = new Array(1, 2, 3); // 创建了数组[1, 2, 3]
var arr2 = [1, 2, 3.14, 'Hello', null, true];
```

越界不报错
``` js
arr1[0]  // 1
arr1[3]  // undefined

// 如果通过索引赋值时，索引超过了范围，统一可以赋值
arr1[3] = 3
arr1[3] // 3
```

#### push和pop

+ push()向Array的末尾添加若干元素
+ pop()则把Array的最后一个元素删除掉

```js
var arr = [1, 2];
arr.push('A', 'B'); // 返回Array新的长度: 4
arr; // [1, 2, 'A', 'B']
arr.pop(); // pop()返回'B'
arr; // [1, 2, 'A']
arr.pop(); arr.pop(); arr.pop(); // 连续pop 3次
arr; // []
arr.pop(); // 空数组继续pop不会报错，而是返回undefined
arr; // []
```

#### unshift和shift

+ unshift()往Array的头部添加若干元素
+ shift()方法则把Array的第一个元素删掉

```js
var arr = [1, 2];
arr.unshift('A', 'B'); // 返回Array新的长度: 4
arr; // ['A', 'B', 1, 2]
arr.shift(); // 'A'
arr; // ['B', 1, 2]
arr.shift(); arr.shift(); arr.shift(); // 连续shift 3次
arr; // []
arr.shift(); // 空数组继续shift不会报错，而是返回undefined
arr; // []
```

#### splice

splice()方法是修改Array的“万能方法”，它可以从指定的索引开始删除若干元素，然后再从该位置添加若干元素

```js
var arr = ['Microsoft', 'Apple', 'Yahoo', 'AOL', 'Excite', 'Oracle'];
// 从索引2开始删除3个元素,然后再添加两个元素:
arr.splice(2, 3, 'Google', 'Facebook'); // 返回删除的元素 ['Yahoo', 'AOL', 'Excite']
arr; // ['Microsoft', 'Apple', 'Google', 'Facebook', 'Oracle']
// 只删除,不添加:
arr.splice(2, 2); // ['Google', 'Facebook']
arr; // ['Microsoft', 'Apple', 'Oracle']
// 只添加,不删除:
arr.splice(2, 0, 'Google', 'Facebook'); // 返回[],因为没有删除任何元素
arr; // ['Microsoft', 'Apple', 'Google', 'Facebook', 'Oracle']
```

#### sort和reverse

+ sort()可以对当前Array进行排序
+ reverse()把整个Array的元素给调个个，也就是反转

```js
var arr = ['B', 'C', 'A'];
arr.sort();
arr; // ['A', 'B', 'C']
arr.reverse();
arr; // ['C', 'B', 'A']
```

#### concat和slice

+ concat()方法把当前的Array和另一个Array连接起来，并返回一个新的Array
+ slice()就是对应String的substring()版本，它截取Array的部分元素，然后返回一个新的Array

```js
var arr = ['A', 'B', 'C'];
var added = arr.concat([1, 2, 3]);
added; // ['A', 'B', 'C', 1, 2, 3]

var arr = ['A', 'B', 'C', 'D', 'E', 'F', 'G'];
arr.slice(0, 3); // 从索引0开始，到索引3结束，但不包括索引3: ['A', 'B', 'C']
arr.slice(3); // 从索引3开始到结束: ['D', 'E', 'F', 'G']

// 如果不给slice()传递任何参数，它就会从头到尾截取所有元素。利用这一点，我们可以很容易地复制一个Array
var aCopy = arr.slice();
aCopy; // ['A', 'B', 'C', 'D', 'E', 'F', 'G']
aCopy === arr; // false
```


#### vue数组
Vue 将被侦听的数组的变更方法进行了包裹，所以它们也将会触发视图更新。这些被包裹过的方法包括：

+ push()
+ pop()
+ shift()
+ unshift()
+ splice()
+ sort()
+ reverse()

### 对象

JavaScript的对象是一种无序的集合数据类型，它由若干键值对组成

```js
obj1 = new Object()
obj2 = {}
```

由于JavaScript的对象是动态类型，你可以自由地给一个对象添加或删除属性

未定义的属性不报错
```js
obj1.a = 1  
obj1.a // 1
obj1.b  // undefined

obj1.b = 2
obj1.b // 2

// 删除b属性
delete obj1.b
delete obj1.b // 删除一个不存在的school属性也不会报错
```

使用hasOwnProperty, 判断对象是否有该属性

```js
obj1.hasOwnProperty('b') // false
```

### null和undefined

null表示一个空的值，而undefined表示值未定义。事实证明，这并没有什么卵用

```js
var a = {a: 1}
a.b  // undefined
a.b = null
a.b  // null
```

### 逻辑运算符

+ &&: 与运算
+ ||: 或运算
+ !: 非运算

### 关系运算符

大于和小于没啥特别的, 要特别注意相等运算符==。JavaScript在设计时，有两种比较运算符：

+ ==: 它会自动转换数据类型再比较，很多时候，会得到非常诡异的结果；
+ ===: 它不会自动转换数据类型，如果数据类型不一致，返回false，如果一致，再比较。

```js
false == 0; // true
false === 0; // false
```

## 变量


### var申明

var：变量提升（无论声明在何处，都会被提至其所在作用于的顶部）

```js
var age = 20
function f1() {console.log(age)}
f1() // 20
```

### 局部变量声明

let：无变量提升（未到let声明时，是无法访问该变量的）

```js
{ let a1 = 20 }
a1 // a1 is not defined
```

### 申明常量

const：无变量提升，声明一个基本类型的时候为常量，不可修改；声明对象可以修改

```js
const c1 = 20
c1 = 30 // Assignment to constant variable
```


## 字符串

JavaScript的字符串就是用''或""括起来的字符表示

```js
str1 = 'str'
str2 = "str"
```

### 字符串转义

使用转义符: \

```js
'I\'m \"OK\"!';
```


### 多行字符串

```js
ml = `这是一个
多行
字符串`; 
// "这是一个\n多行\n字符串"
```

### 字符串模版

格式: 使用``表示的字符串 可以使用${var_name} 来实现变量替换

```js
var name = '小明'
var age = 20
console.log(`你好, ${name}, 你今年${age}岁了！`)// 你好, 小明, 你今年20岁了！
```

### 字符串拼接

直接使用+号

### 常用操作

+ toUpperCase: 把一个字符串全部变为大写
+ toLowerCase: 把一个字符串全部变为小写


## 错误处理

一种是程序写的逻辑不对，导致代码执行异常

```js
var s = null
s.length 
// VM1760:1 Uncaught TypeError: Cannot read property 'length' of null
//     at <anonymous>:1:3
```

如果在一个函数内部发生了错误，它自身没有捕获，错误就会被抛到外层调用函数，如果外层函数也没有捕获，该错误会一直沿着函数调用链向上抛出，直到被JavaScript引擎捕获，代码终止执行


我们可以判断s的合法性, 在保证安全的情况下，使用
```js
if (s !== null) {s.length}
```

也可以捕获异常, 阻断其往上传传递

### try catch

```js
try { s.length } catch (e) {console.log('has error, '+ e)}
// VM2371:1 has error, TypeError: Cannot read property 'length' of null
```

完整的try ... catch ... finally:
```js
try {
    ...
} catch (e) {
    ...
} finally {
    ...
}
```

+ try: 捕获代码块中的异常
+ catch: 出现异常时需要执行的语句块
+ finally: 无论成功还是失败 都需要执行的代码块


常见实用案例:  loading


## 函数


## 方法


### 箭头函数(匿名函数)


### 错误类型

javaScript有一个标准的Error对象表示错误

```js
err = new Error('异常来')
err 
// <!-- Error: 异常来
//     at <anonymous>:1:7 -->

err instanceof Error
// true
```

### 抛出错误

程序也可以主动抛出一个错误，让执行流程直接跳转到catch块。抛出错误使用throw语句

```js
throw new Error('抛出异常')
// VM3447:1 Uncaught Error: 抛出异常
//     at <anonymous>:1:7
// (anonymous) @ VM3447:1
```

## 条件判断

语法格式:
```js
if (condition) {
    ...
} else if (condition) {
    ...
} else {
    ...
}
```

注意条件需要加上括号, 其他和Go语言的if一样:

```js
var age = 20;
if (age >= 6) {
    console.log('teenager');
} else if (age >= 18) {
    console.log('adult');
} else {
    console.log('kid');
}
```

## for 循环

语法格式:

```js
for (初始条件; 判断条件; 修改变量) {
    ...
}
```

注意条件需要加上括号:

```js
var x = 0;
var i;
for (i=1; i<=10000; i++) {
    x = x + i;
}
x; // 50005000
```

### for in (不推荐使用)

for循环的一个变体是for ... in循环，它可以把一个对象的所有属性依次循环出来


遍历对象: 遍历出来的属性是元素的key

```js
var o = {
    name: 'Jack',
    age: 20,
    city: 'Beijing'
};
for (var key in o) {
    console.log(key); // 'name', 'age', 'city'
}
```

遍历数组: 一个Array数组实际上也是一个对象，它的每个元素的索引被视为一个属性

```js
var a = ['A', 'B', 'C'];
for (var i in a) {
    console.log(i); // '0', '1', '2'
    console.log(a[i]); // 'A', 'B', 'C'
}
```

for in 有啥问题? 为啥不推荐使用, 我们看下面一个例子

当我们手动给Array对象添加了额外的属性后，for ... in循环将带来意想不到的意外效果
```js
var a = ['A', 'B', 'C'];
a.name = 'Hello';
for (var x in a) {
    console.log(x); // '0', '1', '2', 'name'
}
```

为什么? 这和for in的遍历机制相关: 遍历对象的属性名称

那如何解决这个问题喃? 答案是 for of


### for of

for ... of循环则完全修复了这些问题，它只循环集合本身的元素

```js
var a = ['A', 'B', 'C'];
a.name = 'Hello';
for (var x of a) {
    console.log(x); // 'A', 'B', 'C'
}
```


但是我们用for of 能遍历对象吗?

```js
var o = {
    name: 'Jack',
    age: 20,
    city: 'Beijing'
};
for (var key of o) {
    console.log(key);
}
// VM2749:6 Uncaught TypeError: o is not iterable
//     at <anonymous>:6:17
```

变通的方法是: 我们可以通过Object提供的方法获取key数组,然后遍历

```js
var o = {
    name: 'Jack',
    age: 20,
    city: 'Beijing'
};
for (var key of Object.keys(o)) {
    console.log(key); // 'name', 'age', 'city'
}
```


### forEach方法

forEach()方法是ES5.1标准引入的, 他也是遍历元素的一种常用手段, 也是能作用于可跌倒对象上, 和for of一样

```js
arr.forEach(function(item) {console.log(item )})
```

当然这还有一种简洁写法

```js
arr.forEach((item) => {console.log(item)})
```

### for循环应用

如果后端返回的数据不满足我们展示的需求, 需要修改，比如vendor想要友好显示，我们可以直接修改数据






