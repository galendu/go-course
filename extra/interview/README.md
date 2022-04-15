# 常见面试问题

1. protobuf文本中编号的作用?
   每个字段都有唯一的一个数字标识符,这些标识符是用来在消息的二进制格式中来标识唯一字段的,一旦使用就不能改变,其中[1,15]占用一个字节,[16,2047]占用两个字节.

2. 控制并发的手段？
    使用`WaitGroup`来跟踪goroutine的工作是否完成,`Wait`方法会阻塞直到goroutine计数变为0.

3. map是并发安全的吗? atomic是如何实现并发安全的?


4. go里边的锁怎么实现?


5. 类型转换会进行内存复制吗?如何才能不进行内存复制?

    [非类型安全指针](https://gitee.com/infraboard/go-course/blob/master/zh-cn/base/unsafe_pointer.md)

6. 内存逃逸?
    golang的变量会携带一组校验数据,用来证明它的整个生命周期是否是在运行是已经确定的,如果通过校验会在栈上分配内存,否则会在堆上分配内存,也说是逃逸了.
    关键点:"栈分配廉价, 堆分配昂贵"
    内存逃逸的情况主要有:
    1. 发送指针或带有指针的值到channel中: 不能确定哪个goroutine会在channel上接收值,在编译时无法确认变量何时释放
    2. 在slice中保存指针或带有指针的值
    3. slice背后的数组被重新分配了.
    4. 在interface上调用方法.


7. 三次握手,为什么是三次?

8. MySQL的事务隔离级别?默认的隔离级别?默认隔离级别会造成数据的脏读还是幻读? 怎么解决?
   MySQL的事务是数据库操作的最小工作单元,这些操作一起向系统提交,要么都执行要么都回滚.
   MySQL的事务级别由高到低分别是:
      - RU: Read Uncommitted,读未提交, 会出现脏读,幻读,不可重复读
      - RC: Read Committed,读已提交, 会出现不可重复读和幻读.
      - RR: Repeatable Read,可重复读,这个是MySQL的默认隔离级别, 会出现幻读
      - Serializable:串行

   脏读,幻读,不可重复读?
   - 脏读: 事务A被事务B影响了,事务A读取的数据被事务B修改了但是没有提交,如果事务B发生了回滚,事务A读取的数据就是无效的,这就是脏读.
   - 不可重复读: 事务A和事务B在开启了事务后,事务A读取了数据,事务B修改了该数据并且提交了,那么事务A就读取不到之前的值了,读取出来的是最新的值.
   - 幻读: 事务A和事务B, 事务A在表中读取数据,事务B向表中insert了一条数据并提交了,事务A再读取的时候结果和之前一样,读取的是数据的旧值.

9. MySQL索引的实现?
   MySQL的索引是使用B+树实现的,B+树的子节点数和元素数是相同的,同时MySQL中阶数m被设计成了1200,正好是一个内存页的整数倍,B+树的前两层是存放在内存中的,第三层到最后一层存放在磁盘中.


10. Go调度模型?执行完当前运行队列的goroutine会有什么动作?(MPG)

    [并发调度模型](https://gitee.com/infraboard/go-course/blob/master/day10/concurrency_mem.md)

11. goroutine有什么优势?

12. goroutine一般什么情况下会阻塞?系统调用

13. 切片的原理?扩容后是如何增长的?

    [Go语言切片](https://gitee.com/infraboard/go-course/blob/master/zh-cn/base/slice.md)
 
14. append是并发安全的吗?

15. Map的实现原理?是并发安全的吗?

    [Go语言Map](https://gitee.com/infraboard/go-course/blob/master/zh-cn/base/map.md)

16. 内存泄漏了解过吗?什么情况下会发生内存泄漏? (定时器会发生内存泄漏)

17. pprof了解过吗?