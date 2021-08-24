# Singleton

### 单例的定义
> Singleton Design Pattern: 一个类只允许创建一个对象，这个类就是单例类，这种设计模式就叫单利设计模式

### 单例的用处
> 从业务概念上，有些数据在系统中只应该保存一份，比较适合设计为单例类。

### 单例的实现

* 饿汉式单例
> 在类加载期间，对应golang init() 方法，就应该将instance静态实例初始化好，所以insance实例的创建是线程安全的，这样的实现方式不支持延迟加载实例

```
$ go test -bench=. -v
go-design-patterns/Chapter02/singleton/hungry on  main [!] via 🐹 v1.16.6
➜ go test  -bench=. -v
=== RUN   TestGetInstance
--- PASS: TestGetInstance (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/chyidl/go-design-patterns/Chapter02/singleton/hungry
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkGetInstanceParallel
BenchmarkGetInstanceParallel-8   	1000000000	         0.2097 ns/op
PASS
ok  	github.com/chyidl/go-design-patterns/Chapter02/singleton/hungry	0.332s
```

* 懒汉式-双重检测
> 双重检测实现方式支持延迟加载，支持高并发的单例实现方式，只要instance被创建之后，在调用getInstance()函数都不会进入到加锁逻辑中，这种实现方式解决懒汉式并发度低的问题

```
go-design-patterns/Chapter02/singleton/lazy on  main [!] via 🐹 v1.16.6
➜ go test -bench=. -v
=== RUN   TestGetInstance
--- PASS: TestGetInstance (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/chyidl/go-design-patterns/Chapter02/singleton/lazy
cpu: Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
BenchmarkGetInstanceParallel
BenchmarkGetInstanceParallel-8   	1000000000	         0.7532 ns/op
PASS
ok  	github.com/chyidl/go-design-patterns/Chapter02/singleton/lazy	1.010s

```
