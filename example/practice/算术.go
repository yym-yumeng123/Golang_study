package main

import "fmt"

func main() {
	// 97天假期, 还有几个星期
	var weeks int8 = 97 / 7
	var day int8 = 97 % 7
	fmt.Println(weeks, day)

	// 定义一个变量保存温度, 华氏 -> 摄氏度 5 / 9 * (华氏 - 100)
	var temp float32 = 124.2
	var ss = 5.0 / 9 * (temp - 100)
	fmt.Println(ss)

	a := 9
	b := 2

	// 定义一个临时变量, 交换值
	t := a
	a = b
	b = t

	// 不使用中间变量
	a = a + b
	b = a - b
	a = a - b

	a1 := 100
	fmt.Println("a1 的地址", &a1)
	var ptr *int = &a1
	fmt.Println("ptr指向", *ptr)

	// 两个数的最大值
	n1 := 10
	n2 := 20
	if n1 > n2 {
		fmt.Println("n1最大")
	} else {
		fmt.Println("n2最大")
	}
}
