```go
package main // 表示一个课独立执行的程序

import "fmt"

/**
注意: { 不能单独放在一行
*/
func main() {
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello, World!")
}
```

1. `package main` 定义了包名. 必须在源文件中非注释的第一行指明这个文件属于哪个包
2. `import "fmt"` 告诉 Go 编译器需要使用 `fmt` 包(的函数, 或其他元素), fmt 包实现了格式化 `IO(输入输出)`的函数
3. `func main()` 程序开始执行的函数. `main`函数是每一个可执行程序所必须包含的, 一般来说是在启动后第一个执行的函数
4. `/*...*/` 是注释
5. `fmt.Println(...)` 可以将字符串输出到控制台，并在最后自动增加换行字符 \n

### 执行 Go 程序

```bash
go run hello.go

# 生成二进制文件
go build hello.go
```

### Go 语言基础语法

1. 行分隔符

在 Go 程序中, 一行代表一个语句结束, 每个语句不需要以分号结尾 `;`, 这些工作由 Go 编辑器自动完成

2. 注释

```go
// 单行注释

/*
  多行注释
  多行
*/
```

3. 标识符

标识符用来命名变量, 类型等程序实体, 一个标识符实际上就是一个或是多个字母(A~Z 和 a~z)数字(0~9)、下划线\_组成的序列

```txt
有效标识符
mahes abc a_123 na04 _temp j a23b2 reVal

无效标识符
1ab(以数字开头)
case 关键字
a+b 运算符不允许
```

4. 字符串连接

```go
fmt.Println("Google" + "Robot")
```

5. Go 语言的空格

在 Go 语言中，空格通常用于分隔标识符、关键字、运算符和表达式，以提高代码的可读性

```go
// 变量的声明必须使用空格隔开
var x int
const Pi float64 = 3.14159265358979323846

// 运算符和操作数之间要使用空格能让程序更易阅读
fruit = apples + oranges;

// 在函数调用时，函数名和左边等号之间要使用空格，参数之间也要使用空格
result := add(2, 3)
```

6. 格式化字符串

Go 语言中使用 `fmt.Sprintf` 或 `fmt.Printf` 格式化字符串并赋值给新串

- Sprintf 根据格式化参数生成格式化的字符串并返回该字符串
- Printf 根据格式化参数生成格式化的字符串并写入标准输出

```go
package main

import "fmt"

func main() {
   // %d 表示整型数字，%s 表示字符串
  var stockcode=123
  var enddate="2020-12-31"
  var url="Code=%d&endDate=%s"
  var target_url=fmt.Sprintf(url,stockcode,enddate)
  fmt.Println(target_url) // Code=123&endDate=2020-12-31
  fmt.Printf(url,stockcode,enddate) // Code=123&endDate=2020-12-31
}
```
