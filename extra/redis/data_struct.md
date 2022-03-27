# Redis数据结构

![](./images/key-value-data-stores.webp)


## 环境准备

```sh
# 使用Docker启动一个Redis服务
docker run -p 6379:6379 -itd --name redis  redis
# 进入Redis的命令交互界面
docker exec -it redis redis-cli
```

## 关于Key


## Value类型

### Strings

![](./images/redis_strings.png)


#### 基本操作

+ set
+ mset
+ get
+ mget

单值操作:
```
127.0.0.1:6379> set mykey somevalue
OK
127.0.0.1:6379> get mykey
"somevalue"
```

多值操作:
```
127.0.0.1:6379> mset a 10 b 20 c 30
OK
127.0.0.1:6379> mget a b c
1) "10"
2) "20"
3) "30"
```

#### 设置过期

在添加key的时候, 我们可以为其添加参数, 其中最常用的就是 ex参数, 控制过期时长

```
127.0.0.1:6379> set key 100 ex 10
OK
127.0.0.1:6379> ttl key
(integer) 8
127.0.0.1:6379> get key
"100"
127.0.0.1:6379> ttl key
(integer) 2
127.0.0.1:6379> get key
"100"
127.0.0.1:6379> ttl key
(integer) -2
127.0.0.1:6379> get key
(nil)
```

其实 set 还支持很多参数
```
set key value [EX seconds|PX milliseconds|EXAT timestamp|PXAT milliseconds-timestamp|KEEPTTL] [NX|XX] [GET]
```

更多操作请参考 [commands for strings](https://redis.io/commands/?group=string)


#### 应用场景

##### 分布式锁

针对 set 还有2个参数: 
+ NX – Only set the key if it does not already exist.
+ XX – Only set the key if it already exist.

应该redis是并发安全的, 所以我们可以基于此来实现分布式锁

```
# set 如果不存在 就添加了一个 5秒过期的key
# 如果key存在就不会有任何操作
127.0.0.1:6379> set lock_key 1 ex 5 nx
OK
127.0.0.1:6379> set lock_key 1 ex 5 nx
(nil)

# 5秒过后
127.0.0.1:6379> set lock_key 1 ex 5 nx
OK
127.0.0.1:6379> set lock_key 1 ex 5 nx
(nil)

# set 添加
127.0.0.1:6379> set lock_key 1 ex 5 nx
OK
127.0.0.1:6379> del lock_key
(integer) 1
```


##### 计算器

除了有简单的SET,GET操作, Redis为了解决分布式系统的计数问题, 专门支持了一些Counter操作:

+ incr:  +1
+ incrby:+n
+ decr:  -1
+ decrby:-n

```
127.0.0.1:6379> set counter 100
OK
127.0.0.1:6379> incr counter
(integer) 101   
127.0.0.1:6379> incr counter
(integer) 102   
127.0.0.1:6379> decr counter
(integer) 101   
127.0.0.1:6379> incrby counter 100
(integer) 201   
127.0.0.1:6379> decrby counter 100
(integer) 101   
127.0.0.1:6379>
```


### Sets

![](./images/redis_sets.png)


### Sorted Sets

![](./images/redis_zset.png)

### Lists

![](./images/list.png)

### Hashes

![](./images/hashes.png)

### Bitmaps

![](./images/bitmaps.png)

### Bitfields

![](./images/bitmaps.png)

### HyperLogLog

![](./images/bitmaps.png)

### Geospatial indexes

![](./images/data-structures-geospatial.webp)

### Streams

![](./images/streams-2.webp)


## 参考

+ [Data Structures](https://redis.com/redis-enterprise/data-structures/)
+ [Data types tutorial](https://redis.io/docs/manual/data-types/data-types-tutorial/)
+ [Redis data types](https://redis.io/docs/manual/data-types/)