package main

import (
	"fmt"
	"strings"
)

var age = test11()

func test11() int {
	fmt.Println("test()")
	return 90
}

// 完成一些初始化工作
func init() {
	fmt.Println("init")
}

// 返回一个函数数据类型 func(int) int
// 返回的是一个匿名函数, 这个匿名函数引用到函数外的 n, 因此这个匿名函数就和n形成了一个整体, 构成闭包

func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n += x
		return n
	}
}

func makeSuffix(suffix string) func(string2 string) string {
	return func(name string) string {
		// name 无后缀, 加上, 反之, 返回原来
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}

		return name
	}
}

func main() {
	fmt.Println("main")

	// 匿名函数

	// 案例: 求两个值的和
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)

	a := func(n1 int, n2 int) int {
		return n1 + n2
	}

	res2 := a(20, 30)

	fmt.Println(res1)
	fmt.Println(res2)

	// 闭包
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))

	g := makeSuffix(".jpg")
	fmt.Println("文件名处理=", g("winter"))
}
