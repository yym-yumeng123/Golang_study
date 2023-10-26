# Golang_study

学习Golang

### 运行 go 代码

```bash
go run hello.go
```

```go
package main // 表明该文件所在包是 main, 每个文件必须属于一个包
import "fmt" // 引入一个包, 包名 fmt, 可以使用该包的函数

/*
 * func 关键字, 表示一个函数
 * main 函数名, 主函数, 程序的入口
*/
func main()  {
  fmt.Println("Hello World")
}
```