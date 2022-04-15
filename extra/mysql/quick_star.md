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



### RIGHT JOIN


### INNER JOIN



### Left Join且不含B



### Right Join且不含A



### Full Join



### Full Join且不含交集
