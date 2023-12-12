### Goroutine 调度器

Goroutine 占用的资源非常小，上节课我们也说过，每个 Goroutine 栈的大小默认是 2KB。
而且，Goroutine 调度的切换也不用陷入（trap）操作系统内核层完成，代价很低。
因此，一个 Go 程序中可以创建成千上万个并发的 Goroutine。
而将这些 Goroutine 按照一定算法放到“CPU”上执行的程序，就被称为  Goroutine 调度器（Goroutine Scheduler），
注意，这里说的“CPU”打了引号。


一个 Go 程序对于操作系统来说只是一个用户层程序，操作系统眼中只有线程，它甚至不知道有一种叫 Goroutine 的事物存在。
所以，Goroutine 的调度全要靠 Go 自己完成
那么，实现 Go 程序内 Goroutine 之间“公平”竞争“CPU”资源的任务，就落到了 Go 运行时（runtime）头上了

可是，在操作系统层面，线程竞争的“CPU”资源是真实的物理 CPU，但在 Go 程序层面，各个 Goroutine 要竞争的“CPU”资源又是什么呢？

Goroutine 们要竞争的“CPU”资源就是操作系统线程。这样，
Goroutine 调度器的任务也就明确了：`将 Goroutine 按照一定算法放到不同的操作系统线程中去执行`


### Goroutine 调度器模型与演化过程
#### 最初的 G-M 模型

每个 Goroutine 对应于运行时中的一个抽象结构：`G(Goroutine)`

被视作“物理 CPU”的操作系统线程，则被抽象为另外一个结构：`M(machine)`

调度器的工作就是将 G 调度到 M 上去运行。为了更好地控制程序中活跃的 M 的数量，
调度器引入了 `GOMAXPROCS` 变量来表示 Go 调度器可见的“处理器”的最大数量

G-M 模型的一个`重要不足`：限制了 Go 并发程序的伸缩性，尤其是对那些有高吞吐或并行计算需求的服务程序。

- 单一全局互斥锁(Sched.Lock) 和集中状态存储的存在，导致所有 Goroutine 相关操作，比如创建、重新调度等，都要上锁；
- Goroutine 传递问题：M 经常在 M 之间传递“可运行”的 Goroutine，这导致调度延迟增大，也增加了额外的性能损耗；
- 每个 M 都做内存缓存，导致内存占用过高，数据局部性较差
- 由于系统调用（syscall）而形成的频繁的工作线程阻塞和解除阻塞，导致额外的性能损耗

**G-P-M调度模型**

P 是一个“`逻辑 Proccessor`”，每个 G（Goroutine）要想真正运行起来，
首先需要被分配一个 P，也就是进入到 P 的本地运行队列（local runq）中

对于 G 来说，P 就是运行它的“CPU”，可以说：在 G 的眼里只有 P

但从 Go 调度器的视角来看，真正的“CPU”是 M，只有将 P 和 M 绑定，才能让 P 的 runq 中的 G 真正运行起来。

#### 深入 G-P-M 模型

- G: 代表 Goroutine，存储了 Goroutine 的执行栈信息、Goroutine 状态以及 Goroutine 的任务函数等，而且 G 对象是可以重用的；
- P: 代表逻辑 processor，P 的数量决定了系统内最大可并行的 G 的数量，P 的最大作用还是其拥有的各种 G 对象队列、链表、一些缓存和状态；
- M: M 代表着真正的执行计算资源在绑定有效的 P 后，进入一个调度循环，而调度循环的机制大致是从 P 的本地运行队列以及全局队列中获取 G，切换到 G 的执行栈上并执行 G 的函数，调用 goexit 做清理工作并回到 M，如此反复。M 并不保留 G 状态，这是 G 可以跨 M 调度的基础。

### Channel

channel 是`用于 Goroutine 间通信`的，
所以绝大多数对 channel 的读写都被分别放在了不同的 Goroutine 中

```go
// 创建 channel
// 声明了一个元素为 int 类型的 channel 类型变量 ch
var ch chan int // nil


ch1 := make(chan int) // 无缓冲 channel
ch2 := make(chan int, 5) // 带缓冲 channel
```

```go
// 发送与接收
ch1 <- 13    // 将整型字面值13发送到无缓冲channel类型变量ch1中
n := <- ch1  // 从无缓冲channel类型变量ch1中接收一个整型值存储到整型变量n中
ch2 <- 17    // 将整型字面值17发送到带缓冲channel类型变量ch2中
m := <- ch2  // 从带缓冲channel类型变量ch2中接收一个整型值存储到整型变量m中
```

- 对无缓冲 channel 类型的发送与接收操作，一定要放在两个不同的 Goroutine 中进行，否则会导致 deadlock
- Goroutine 对无缓冲 channel 的接收和发送操作是同步的

```go
func main() {
	// 对同一个无缓冲 channel
	// 只有对它进行接收操作的 Goroutine 和对它进行发送操作的 Goroutine 都存在的情况下
	// 通信才能得以进行，否则单方面的操作会让对应的 Goroutine 陷入挂起状态
    ch1 := make(chan int)
    ch1 <- 13 // fatal error: all goroutines are asleep - deadlock!
    n := <-ch1
    println(n)
}
```

- 带缓冲 channel 的发送操作在缓冲区未满、接收操作在缓冲区非空的情况下是异步的
- 对一个带缓冲 channel 来说，在缓冲区未满的情况下，对它进行发送操作的 Goroutine 并不会阻塞挂起；在缓冲区有数据的情况下，对它进行接收操作的 Goroutine 也不会阻塞挂起

```go
ch2 := make(chan int, 1)
n := <-ch2 // 由于此时ch2的缓冲区中无数据，因此对其进行接收操作将导致goroutine挂起

ch3 := make(chan int, 1)
ch3 <- 17  // 向ch3发送一个整型数17
ch3 <- 27  // 由于此时ch3中缓冲区已满，再向ch3发送数据也将导致goroutine挂起
```


```go
ch1 := make(chan<- int, 1) // 只发送channel类型
ch2 := make(<-chan int, 1) // 只接收channel类型

<-ch1       // invalid operation: <-ch1 (receive from send-only type chan<- int)
ch2 <- 13   // invalid operation: ch2 <- 13 (send to receive-only type <-chan int)
```

```go
// 关闭 channel
n := <- ch      // 当ch被关闭后，n将被赋值为ch元素类型的零值
m, ok := <-ch   // 当ch被关闭后，m将被赋值为ch元素类型的零值, ok值为false
for v := range ch { // 当ch被关闭后，for range循环结束
    ... ...
}
```

**select**

当涉及同时对多个 channel 进行操作时，
我们会结合 Go 为 CSP 并发模型提供的另外一个原语 `select`，一起使用。

```go
select {
case x := <-ch1:     // 从channel ch1接收数据
  ... ...

case y, ok := <-ch2: // 从channel ch2接收数据，并根据ok值判断ch2是否已经关闭
  ... ...

case ch3 <- z:       // 将z值发送到channel ch3中:
  ... ...

default:             // 当上面case中的channel通信均无法实施时，执行该默认分支
}
```

无缓冲 channel 的惯用法

- 用作信号传递
- 用于替代锁机制

带缓冲 channel 的惯用法

- 用作消息队列
- 用作计数信号量（counting semaphore）