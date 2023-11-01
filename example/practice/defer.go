package main

import "fmt"

func sum(n1 int, n2 int) int {
	// 当执行到 defer, 暂时不执行, 会将 defer 后面的语压入到独立的栈 (defer栈) 先进后出
	// 函数执行完毕, 从 defer 栈, 执行
	defer fmt.Println("n1=", n1) // 3. n1= 10
	defer fmt.Println("n2=", n2) // 2. n2= 20

	// 增加
	n1++
	n2++

	res := n1 + n2
	fmt.Println("res=", res) // 1. res= 30
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res) // 4. res= 30
}
