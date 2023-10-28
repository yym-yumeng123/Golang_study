package main

import "fmt"

func main() {
	// 基本数据类型在内存布局
	var i int = 10
	// i 的地址 &i
	fmt.Println("i的地址", &i)

	// 指针类型
	// ptr 是一个指针变量; ptr 类型是 *int; ptr本身的值 &i
	var ptr *int = &i
	fmt.Printf("ptr=%v", ptr)          // 0xc00000a0b8
	fmt.Printf("ptr 的地址= %v", &ptr)    // 0xc00004e028
	fmt.Printf("ptr 指向的值= %v\n", *ptr) // 10

	// 获取整数地址
	var num int = 100
	fmt.Println("num 地址= \n", &num)

	var ptr1 *int
	// ptr1 获取 &num 的地址
	ptr1 = &num
	// *ptr 获取指针类型指向的值
	*ptr1 = 10
	fmt.Println("ptr1=", num, ptr1)
}
