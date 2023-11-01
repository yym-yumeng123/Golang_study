package main

import "fmt"

func cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符好错误")
	}

	return res
}

func test(n3 int) {
	n3 = n3 + 1
	fmt.Println("test n3=", n3)
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFunc(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 可变参数
func sum1(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}

func main() {
	var n1 float64 = 1.2
	var n2 float64 = 2.3
	var operator byte = '-'
	result := cal(n1, n2, operator)
	fmt.Println("res=", result)

	// 调用test
	n3 := 10
	test(n3)                    // 11
	fmt.Println("main n3=", n3) // 10

	// 函数也是一种数据类型, 赋值给另一个变量, 该变量可以对函数调用
	a := getSum
	fmt.Printf("a 的类型 %T, getSum类型是%T", a, getSum) // func(int, int) int
	res := a(10, 40)
	fmt.Println(res)

	res2 := myFunc(getSum, 50, 50)
	fmt.Println("res2=", res2)

	res4 := sum1(1, 2, 3, 4, 5)
	fmt.Println("res4=", res4)
}
