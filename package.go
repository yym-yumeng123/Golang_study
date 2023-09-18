package main

import "fmt" // 包名

func init() {
	// 在 main 方法之前调用
	fmt.Println("我打印了 init")
}

func main() {
	fmt.Println("我是 main")
}
