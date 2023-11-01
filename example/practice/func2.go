package main

import "fmt"

func test1(n int) {
	if n > 2 {
		n--
		test1(n)
	}
	fmt.Println("n=", n)
}

func fibonaci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fibonaci(n-1) + fibonaci(n-2)
	}
}

func fn1(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*fn1(n-1) + 1
	}
}

func peach(n int) int {
	if n > 10 || n < 1 {
		fmt.Println("天数不对")
		return 0
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}

func main() {
	// 递归, 先入栈, 后出栈 (先入后出)
	// n= 2
	// n= 2
	// n= 3
	test1(4)

	// 斐波那契数 1 1 2 3 5 8 13 ..., 给一个整数, 求出斐波那契数
	fmt.Println(fibonaci(4))
	fmt.Println(fibonaci(5))
	fmt.Println(fibonaci(6))

	// f(1)=3 f(n)=2*f(n-1)+1, 使用递归思想
	fmt.Println(fn1(6))

	// 有一堆桃子, 猴子第一天吃了其中一半, 并多吃了一个, 以后每天猴子都吃其中一半, 然后多吃一个, 当第十天, 想再吃(还没吃,) 只有一个了, 最初有多少桃子
	// 反推, 第10天, 1个桃子
	// 第9天 , 桃子 = (第10天桃子 + 1) *2
	// 第 n 天, peach(n) = (peach(n+1) + 1) * 2
	fmt.Println(peach(1))
	fmt.Println(peach(9))
}
