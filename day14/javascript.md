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


### null和undefined

null表示一个空的值，而undefined表示值未定义。事实证明，这并没有什么卵用

```js
var a = {a: 1}
a.b  // undefined
a.b = null
a.b  // null
```

### 数组

```js
var arr1 = new Array(1, 2, 3); // 创建了数组[1, 2, 3]
var arr2 = [1, 2, 3.14, 'Hello', null, true];
```

越界不报错
``` js
arr1[0]  // 1
arr1[3]  // undefined
```


### 对象

```js
obj1 = new Object()
obj2 = {}
```

未定义的属性不报错
```js
obj1.a = 1  
obj1.a // 1
obj1.b  // undefined
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

