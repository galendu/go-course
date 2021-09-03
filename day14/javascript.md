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

+ var：变量提升（无论声明在何处，都会被提至其所在作用于的顶部）
+ let：无变量提升（未到let声明时，是无法访问该变量的）
+ const：无变量提升，声明一个基本类型的时候为常量，不可修改；声明对象可以修改



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