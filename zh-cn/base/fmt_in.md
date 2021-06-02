# 字符串格式化

Go语言用于控制文本输入和输出格式的库是fmt

## 格式化输入

fmt包中提供了3类读取输入的函数：

+ Scan家族：从标准输入os.Stdin中读取数据，包括Scan()、Scanf()、Scanln()
+ SScan家族：从字符串中读取数据，包括Sscan()、Sscanf()、Sscanln()
+ Fscan家族：从io.Reader中读取数据，包括Fscan()、Fscanf()、Fscanln()

其中:

+ Scanln、Sscanln、Fscanln在遇到换行符的时候停止
+ Scan、Sscan、Fscan将换行符当作空格处理
+ Scanf、Sscanf、Fscanf根据给定的format格式读取，就像Printf一样

这3家族的函数都返回读取的记录数量，并会设置报错信息，例如读取的记录数量不足、超出或者类型转换失败等

以下是他们的定义:

```sh
go doc fmt | grep -Ei "func [FS]*Scan"
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
func Scan(a ...interface{}) (n int, err error)
func Scanf(format string, a ...interface{}) (n int, err error)
func Scanln(a ...interface{}) (n int, err error)
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
```

因为还没介绍io.Reader，所以Fscan家族的函数暂且略过，但用法和另外两家族的scan类函数是一样的

## 作业

问题搜索工具, 