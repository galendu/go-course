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

#### 数据库名

数据库整体命名规则采用snak缩写规则, 比如 cmdb_service ,注意禁止使用中横线 比如cmdb-service

其他情况按照报错提醒 去掉不合法字符就可以

#### 字符集

字符集字面理解就是 一堆字符的合集, 每个字符集包含的字符个数不同, 比如:
+ 英文字符集: ASCII, 采用7位编码, 总共能编码2^7个字符
+ 中文字符集(兼容ASCII): GB2312、BIG5、GBK, GB18030, 都采用双字节字进行编码,
+ 万国码(兼容ASCII): Unicode

#### 排序规则


### 定义表


#### 表结构(字段)



#### 索引



#### 完整性约束



## DML(数据操作语句)



## 联表查询

![](./images/sql_join.jpeg)

在进行关联查询之前 我们需要至少准备2张表（现实中的项目往往比较复杂, 5，6张表联合查询是常事儿）

我们以用户系统为例:

+ 用户表: t_user

![](./images/t_user.png)
```
mysql> select * from t_user;
+----+--------+---------------+
| id | name   | department_id |
+----+--------+---------------+
|  1 | 张三 |             1 |
|  2 | 王五 |             0 |
+----+--------+---------------+
```



+ 部门表: t_department

![](./images/t_department.png)
```
mysql> select * from t_department;
+----+-----------+
| id | name      |
+----+-----------+
|  1 | 市场部 |
|  3 | 研发部 |
+----+-----------+
```


### LEFT JOIN

![](./images/left_join.webp)

以左表为准, 把符合条件的关联过来, 如果没有则使用null

比如查询用户的同时，查询出用户所属的部门
```sql
SELECT
	u.*,
	d.name 
FROM
	t_user u
	LEFT JOIN t_department d ON u.department_id = d.id

-- ON 也可以添加多个条件
```

![](./images/left_join_exm.png)

注意:
+ department 1 右表有数据
+ department 0 右表无数据
+ department 3 左表无数据

### RIGHT JOIN

![](./images/right_join.webp)

以右表为准, 把符合条件的关联过来, 如果没有则使用null

```sql
SELECT
	u.*,
	d.name 
FROM
	t_user u
	RIGHT JOIN t_department d ON u.department_id = d.id
```

![](./images/right_join_exm.png)

注意:
+ 张三 部门1        左边表有数据
+ 王五 部门0        不符合关联条件 无数据
+ 市场部            右表有, 左表无数据

### INNER JOIN

![](./images/inner_join.webp)

意思就是取交集，就是要两边都有的东西，所以也就是不能有null出现

```sql
SELECT
	u.*,
	d.name 
FROM
	t_user u
	INNER JOIN t_department d ON u.department_id = d.id
```

![](./images/inner_join_exm.png)

### Left Join且不含B

![](./images/left_join_not_b.webp)

A中与B没有交集的部分，所以就是，join B表会得到null的内容, 比如获取哪些用户没有部门

```sql
SELECT
	u.*,
	d.name 
FROM
	t_user u
	LEFT JOIN t_department d ON u.department_id = d.id WHERE d.name is NULL
```

![](./images/left_join_2.png)


### Right Join且不含A

![](./images/right_join_not_a.webp)

同理，就是与上面的情况相反, 不如我们需要筛选出那么部门没有人

```sql
SELECT
	u.*,
	d.name 
FROM
	t_user u
	RIGHT JOIN t_department d ON u.department_id = d.id WHERE u.id IS NULL
```


![](./images/right_join_2.png)



### Full Join

![](./images/full_join.webp)

mysql语法不支持full outer join, 也就是说我们无法通过一个语句来实现集合的求和操作, 所以我们用union来实现, UNION 操作符用于合并两个或多个 SELECT 语句的结果集

union 是对数据进行并集操作, 因此需要要求数据: 
+ 合并集合的结果有相同个数的列
+ 并且每个列的类型是一样的

1. Union: 合并集合, 并且数据去重
```sql
SELECT
	u.*,
	d.NAME 
FROM
	t_user u
	LEFT JOIN t_department d ON u.department_id = d.id 
UNION
SELECT
	u.*,
	d.NAME 
FROM
	t_user u
	RIGHT JOIN t_department d ON u.department_id = d.id
```

![](./images/union_1.png)

2. Union ALL: 合并集合, 不去重
```sql
SELECT
	u.*,
	d.NAME 
FROM
	t_user u
	LEFT JOIN t_department d ON u.department_id = d.id 
UNION ALL
SELECT
	u.*,
	d.NAME 
FROM
	t_user u
	RIGHT JOIN t_department d ON u.department_id = d.id
```

![](./images/union_2.png)

因为Union ALL 少了去重的操作, 性能上会把Union好很多, 特别是当集合特别大的时候

3. union all 自己计算

比如我们把红色部分拆分为2部分:
+ A和B的左连接(A + AB公共的)
+ A和B的右连接去除公共部分(B独有的部分)

```sql
SELECT
	u.*,
	d.NAME 
FROM
	t_user u
	LEFT JOIN t_department d ON u.department_id = d.id 
UNION ALL
SELECT
	u.*,
	d.NAME 
FROM
	t_user u
	RIGHT JOIN t_department d ON u.department_id = d.id WHERE u.id IS NULL
```

![](./images/union_3.png)


### Full Join且不含交集

![](./images/full_join_not.webp)

还是用union来实现full outer join

```sql
SELECT
	u.*,
	d.NAME 
FROM
	t_user u
	LEFT JOIN t_department d ON u.department_id = d.id WHERE d.id IS NULL 
UNION ALL
SELECT
	u.*,
	d.NAME 
FROM
	t_user u
	RIGHT JOIN t_department d ON u.department_id = d.id WHERE u.id IS NULL
```

![](./images/union_4.png)

## 子查询



## 常用函数



### 字符函数

+ substr
+ length
+ contact


### 日期函数

+ NOW()
+ UNIX_TIMESTAMP


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