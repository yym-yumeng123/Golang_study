package main

import "fmt"

func main() {
	// 基本数据类型转换
	var i int = 100
	// 把 i 类型转换为 float32
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)

	fmt.Printf("i=%v, n1=%v, n2=%v", i, n1, n2)
}
