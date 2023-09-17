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


































