# MySQL快速起步


## 安装

基于Docker安装:
```sh
$ docker run -p 3306:3306 -itd -e MARIADB_USER=cmdb -e MARIADB_PASSWORD=123456 -e MARIADB_ROOT_PASSWORD=123456 --name mysql   mariadb:latest
```

## DDL(数据定义语句)

Database Define Language缩写, 也就是用于创建数据库和表的SQL语法

### 定义库

![](./images/create_table.png)
 

### 定义表


### 表结构



### 数据类型



### 完整性约束



## DML(数据操作语句)



## 联表查询

![](./images/sql_join.jpeg)


### LEFT JOIN

![](./images/left_join.webp)


### RIGHT JOIN

![](./images/right_join.webp)


### INNER JOIN

![](./images/inner_join.webp)


### Left Join且不含B

![](./images/left_join_not_b.webp)


### Right Join且不含A

![](./images/right_join_not_a.webp)


### Full Join

![](./images/full_join.webp)

### Full Join且不含交集

![](./images/full_join_not.webp)



## 子查询



## 常用函数



### 字符函数

+ substr
+ length
+ contact


### 日期函数

+ NOW()
+ TIMESTAMP


## 常用语句


### DISTINCT


### ON DUMPLICATE KEY


### SELECT INTO


### GROUP_CONTAT


### IF与IFNULL


### CASE语句


### DELETE 联表



## 参考

+ [MySQL字符集和排序规则](https://segmentfault.com/a/1190000020339810)
+ [MySQL 中的三中循环 while loop repeat 的基本用法](https://www.cnblogs.com/Luouy/p/7301360.html)
+ [MySQL里面的子查询的基本使用](http://www.codebaoku.com/it-mysql/it-mysql-218378.html)
+ [MySQL 子查询优化](https://www.jianshu.com/p/3989222f7084)
+ [MySQL—基于规则优化 子查询优化](https://www.rsthe.com/archives/mysql%E5%9F%BA%E4%BA%8E%E8%A7%84%E5%88%99%E4%BC%98%E5%8C%96%E5%AD%90%E6%9F%A5%E8%AF%A2%E4%BC%98%E5%8C%96)