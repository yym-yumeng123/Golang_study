channel 是一种用于在不同 `goroutine` 之间进行通信的原语

```go
ch := make(chan int)
ch := make(chan int, 10)
ch <- 1
fmt.Println(<-ch)

close(ch)
```

#### 关闭一个已经关闭的 `channel`: panic

```go
// panic: close of closed channel
func main() {
  ch := make(chan int, 10)
	close(ch)
	close(ch)
}
```

#### 从关闭的 channel中读取数据: 如果 channel 中有数据, 会继续读, 没有数据会读出零值

```go
func main() {
	ch := make(chan int, 10)
	//ch <- 0
	close(ch)

	num, ok := <-ch
	fmt.Println(ok)  // false
	fmt.Println(num) // 0
}
```

#### 向一个已经关闭的channel中写数据: `panic: send on closed channel`


#### 使用for循环读取已经关闭的channel，会读出channel中的数据然后结束

```go

func main() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	close(ch)

	for num := range ch {
		fmt.Printf("num: %d\n", num)
	}
}
```

#### 从一个未初始化的 channel 中读取数据都会被永远堵塞 `deadlock`



---

### Select 使用场景

`select` 用于同时等待多个 `channel` 中的读写操作

```go
select {
  case <- channel1:
		// 当channel1 可以读取时执行的代码
  case data := <-channel2:
		// 当 channel2可以读取执行的代码, 将读取的数据存储在 data 变量中
  case channel3 <- data:
		// 当向 channel3 发送数据时执行的代码
  default:
		// 当没有任何通道准备好时执行的代码
}
```

select 语句会等待多个 channel 中的任意一个可以进行读写操作，并执行对应的分支语句

如果有多个 channel 同时可以进行读写操作，select 会随机选择一个进行处理

如果所有的 case 子句都无法执行，并且没有 default 子句，select 语句会一直阻塞直到其中一个 case 子句可以执行为止

**需要注意的问题:**

1. select 语句会等待多个 channel 中的任意一个可以进行读写操作
2. 不能在 `select` 语句中使用相同的 `unbuffer channel` 进行读写操作
3. 如果一个channel已经被关闭了, 再向它写入数据会导致 panic

