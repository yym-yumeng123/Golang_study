package main

import "fmt"

func add1(x, y int) (ans int) {
	ans = x + y
	return ans
}

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	i, j := swap(1, 2)
	println(add1(0, 0))
	println(swap(0, 0))
	println(i, j)
	fmt.Println(myFunc(1, 2, 3, 4, 5))

	// 匿名函数
	fmt.Println(func(a, b int) int {
		return a + b
	}(1, 2)) // 3

	println(addSum(1))
	println(addSum(2))
	println(addSum(3))
	println(sum)

	// 使用闭包
	add1 := addFunc()
	println(add1(10))
	println(add1(20))
	println(add1(30))

}

// 闭包使用 sum
var sum = 0

func addSum(x int) int {
	sum += x
	return sum
}

// 可变参数 ...int
func myFunc(args ...int) (ans int) {
	for _, arg := range args {
		ans += arg
	}
	return ans
}

func addFunc() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
