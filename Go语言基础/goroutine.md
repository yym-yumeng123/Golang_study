### Goroutine的概念和使用方法

goroutine 是一种轻量级的线程实现, 可以在单个线程中并发的运行多个 `goroutine`

启动非常简单, 只需要在函数或者方法调用钱价格关键字`go`即可

```go
go func() {
	// goroutine body
}
```

#### 注意

在使用 goroutine 时, 如果忘记等待 goroutine 结束, 可能会导致程序在执行完毕前就退出,\
从而导致 `goroutine` 被取消

```go
func someFunc() {
	time.Sleep(5 * time.Second)
	fmt.Println("some Func finished")
}

func main() {
	go someFunc()
	fmt.Println("main finished")
}
```

### 竞态条件

```go
package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
```

在上面的例子中, 启动了 1000 个 goroutine, 每个 goroutine 都会将计数器变量加 1. 
由于多个 goroutine 同时访问该变量, 因此可能会出现竞态条件

由于多个 goroutine 同时修改 counter 变量, 程序的输出结果可能不确定, 每次运行结果都不同,

要避免这种竞态结果, 可以使用 Go 语言中的锁`sync.Mutex`或`通道机制`来进行同步