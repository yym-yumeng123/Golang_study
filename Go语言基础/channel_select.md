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

