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
}
