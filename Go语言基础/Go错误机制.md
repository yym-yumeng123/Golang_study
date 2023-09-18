### Golang 错误处理概述

Go中, 错误处理机制是通过 返回值和错误类型实现的, 而不是使用异常.

函数通常会返回一个值和一个错误类型, 其中错误类型是一个预定义的接口类型

```go
type error interface {
	Error() string
}
```

如果函数执行成功, 则返回值为期望的结果, 而错误类型为 nil, 表示没有错误发生. 如果函数执行失败,
则返回一个默认值(如果有的话), 而错误类型为一个非 nil 的错误对象.

```go
func add4(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a < 0 || b < 0 ")
	}
	return a + b, nil
}
```

格式化 error 创建

```go
func add4(a, b int) (int, error) {
	if a < 0 || b < 0 {
		// 格式化
		return 0, fmt.Errorf("a < 0 || b < 0, a:%v, b: %v", a, b)
	}
	return a + b, nil
}
```

#### 自定义错误

```go
type MyError struct {
	message string
	code int
}

func (e MyError) Error() string {
	return e.message
}

func IsMyError(err error) bool {
	_, ok := err.(MyError)
	return ok
}

io.EOF
```

### panic 与 recover

Golang 没有传统意义上的异常机制, 而是采用了 `panic` 和 `recover` 机制来处理 `运行时`
错误.

当程序遇到无法处理的错误或者异常情况时, 会使用 `panic` 函数来抛出一个 `panic` 异常.
`panic` 函数接受一个任意类型的参数作为错误信息, 一旦执行 panic 函数, 程序会`立即停止`当前函数的执行,
并向调用者抛出 `panic` 异常. 如果没有被捕获, 该异常会一直沿着调用栈向上抛出, 知道被程序中止或者被`recover` 捕获

```go
type any = interface{}
func panic(v any)
func recover() any
```

`recover` 函数用于捕获 `panic` 异常, `recover`函数必须在 defer 函数中调用, 并且必须在可能会触发 `panic` 异常的代码块中使用
当程序调用 recover 函数时, 如果当前代码块中发生了 `panic` 异常, recover 函数会返回该异常的值
并且程序可以继续执行. 没有发生 panic 异常, recover 函数返回 `nil`


#### 什么情况下使用 panic & 最佳实践

- 尽量避免使用 panic
- 在 panic 之前, 应该尽可能记录信息, 多打log
- 不要在 defer 中使用 panic, 如果函数中有多个defer, 并且其中一个 defer 使用了panic, 那么所有的 defer 都将被执行, 但是 panic 将在其他 defer 执行完成后才会被触发

```go
func main() {
  defer func() {
		println(111)
}
defer func() {
  panic("1")
}
defer func() {
  println(222)
}
}
```



1. 初始化使用
```go
func main(){
	db, err:= sql.Open("mysql", "user:password...")
	if err != nil {
		panic(fmt.Errorf("无法连接数据库:%w", err))
}
	defer db.Close()
	// ...
}
```

---

### errors 包

```go
import "errors"

func doSomething() error {
  return errors.New("something went wrong")
}
```

1. `is`




















