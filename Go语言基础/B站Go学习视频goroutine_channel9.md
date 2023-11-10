### goroutine 协程

**进程和线程说明:**

1. 进程就是程序在操作系统中的一次执行过程, 是系统进行资源分配和调度的基本单位
2. 线程是进程的一个执行实例, 是程序执行的最小单元, 它是比进程更小的能独立运行的基本单位
3. 一个进程可以创建销毁多个线程, 同一个进程中的多个线程可以并发执行
4. 一个程序至少有一个进程, 一个进程至少有一个线程

---

**并发和并行:**

1. 多线程程序在单核上运行, 就是并发
   - 因为是在一个 cpu 上,比如有10个线程, 每个线程执行10ms, 在微观角度看, 在一个时间点上, 只有一个任务在执行
2. 多线程程序在多核上运行, 就是并行
   - 因为是在多个 cpu 上, 比如有10个线程, 每个线程执行10ms, 从微观角度看, 在同一时间点, 有10个线程在执行, 这就是并行

---

**Go协程和Go主线程**

1. Go主线程(由程序员直接称为线程/进程): 一个Go线程, 可以起多个协程, 你可以理解, 协程是轻量级的线程
2. Go 协程 (goroutine) 的特点
   - 有独立的栈空间
   - 共享程序的堆空间
   - 调度由用户控制
   - 协程是轻量级的线程

---

**goroutine 入门小结**

1. 主线程是一个物理线程, 直接作用在 cpu 上, 是重量级的, 非常耗费cpu资源.
2. 协程从主线程开启的, 是轻量级的线程, 是逻辑态. 对资源消耗相对小
3. Golang的协程机制是重要的特点, 可以轻松开启上万个协程, 


**goroutine调度模型**

MPG模式基本介绍

Go的调度器有三个重要的结构: M P G 

1. M: 操作系统的主线程(物理线程)
   - 代表真正的内核OS线程, 真正干活的人
2. P: 协程执行需要的上下文
   - 可以看做一个局部的调度器, 使go代码在一个线程上跑, 是实现 N:1, N:N 映射的关键
3. G: 协程
   - 代表一个goroutine, 它有自己的栈, 用于调度


---

**设置Golang运行的cpu数**

为了充分利用多cpu的优势, 在Golang程序中, 设置运行的cpu数目

- `NumCPU()` 返回本地机器的逻辑CPU个数
- `GOMAXPROCS(n int)` 设置课同时执行的最大CPU数, 并返回先前的设置, n < 1, 就不会更改当前设置



### channel 管道

看个问题

```go
var (
	myMap = make(map[int]int, 10)
)

// test 函数就是计算 n!, 将这个结果放入到 myMap
func test(n int) {
	res := 1
	for i := 0; i <= n; i++ {
		res *= i
	}

	myMap[n] = res
}

func main() {
	// 开启多个协程
	for i := 0; i < 201; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10)

	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
}

```

不同 goroutine 之间如何通信

1. 全部变量加锁同步
2. channel

使用全局变量加锁同步改进程序

- 因为没有对全局变量 m 加锁, 会出现资源争夺文艺, 代码会出现错误, 提示 `concurrent map writes`
- 解决方案, 加入互斥锁
- 我们的阶乘很大,结果会越界, 将求阶乘改为 sum += unit64(i)

```go
// 加锁
var (
	myMap = make(map[int]int, 10)
	// 声明一个全局互斥锁
	// sync 同步
	// Mutex 互斥
	lock sync.Mutex
)

// test 函数就是计算 n!, 将这个结果放入到 myMap
func test(n int) {
	res := 1
	for i := 0; i <= n; i++ {
		res *= i
	}
	// 加锁
	lock.Lock()
	myMap[n] = res
	// 解锁
	lock.Unlock()
}

func main() {
	// 开启多个协程
	for i := 0; i < 20; i++ {
		go test(i)
	}

	// 主线程休眠 10 s, 是让协程完成任务
	// 如果没有这句话, 主线程很快就退出, m 中还没有结果
	time.Sleep(time.Second * 5)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
```

前面使用全局加锁来解决 goroutine 通讯,不完美

1. 主线程在等待所有 goroutine 全部完成的时间还难确
2. 如果主线程休眠时间长了, 会加长时间等待, 如果等待时间短了, goroutine还在工作, 也会随着主线程的退出而销毁
3. 通过全局变量加锁同步来实现通讯, 并不利于多个协程对全局变量的读写操作
4. 一个新的通讯机制 - channel

Channel管道

1. channel 的本质就是队列, 先进先出
2. 线程安全, 多 goroutine 访问时, 不需要加锁, channel 本身就是线程安全的
3. channel是有类型的, 一个 string 类型的 channel 只能放 string 类型

```go
// 基本使用
var 变量名 chan 数据类型

var intChan chan int (intChan 用于存放 int 类型)
var mapChan chan map[int]string (用于存放 map[int]string)
var perchan chan Person 
var perChan2 chan *Person

// 初始化
intChan = make(chan int, 3)
// 写入数据
intChan <- 11

// 读取数据
num1 = <-intChan

<-intChan
```
- channel 是引用类型
- channel必须初始化才能写入数据, 即 make后才能使用
- channel 是有类型的, intChan 只能写入整数int
- channel 的数据放满以后,不能再放入了
- channel 的数据取完了, 不能再取了


channel 的遍历和关闭

使用内置函数 `close` 可以关闭 channel, 当channel关闭后, 就不能再向channel写输了, 但是仍可以读数据

channel 支持 `for-range` 方式遍历

1. 在遍历时, 如果没有关闭 channel, 则会出现 deadlock 错误
2. 遍历时, 如果 channel 已经关闭, 则会正常遍历数据

**阻塞**

读快写慢, 写慢就会阻塞