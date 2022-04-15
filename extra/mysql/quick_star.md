# MySQL快速起步


## 安装

基于Docker安装:
```sh
$ docker run -p 3306:3306 -itd -e MARIADB_USER=cmdb -e MARIADB_PASSWORD=123456 -e MARIADB_ROOT_PASSWORD=123456 --name mysql   mariadb:latest
```

## DDL(数据定义语句)


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


## 常用语言


### SELECT INTO




## 参考

+ [MySQL里面的子查询的基本使用](http://www.codebaoku.com/it-mysql/it-mysql-218378.html)
+ [MySQL 子查询优化](https://www.jianshu.com/p/3989222f7084)
+ [MySQL—基于规则优化 子查询优化](https://www.rsthe.com/archives/mysql%E5%9F%BA%E4%BA%8E%E8%A7%84%E5%88%99%E4%BC%98%E5%8C%96%E5%AD%90%E6%9F%A5%E8%AF%A2%E4%BC%98%E5%8C%96)