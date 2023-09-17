package main

import (
	"fmt"
	"os"
)

func f() {
	// 在函数返回前执行某些操作
	defer func() {
		fmt.Println("f defer") // 3
	}()

	// 先进后出, 后进先出, 栈
	defer func() {
		fmt.Println("第几次")
	}()

	fmt.Println("in f") // 2
}

func main() {
	fmt.Println("call f before") // 1
	f()
	fmt.Println("call f after") // 4

	// defer 语句中函数参数的值在 defer 语句被定义时就已经确定了, 而不是在 defer 语句被执行时才确定
	num := 1
	defer func(x int) {
		println("defer", x)
	}(num)
	num = 100
}

func something() {
	file, err := os.Create("test.txt")
	if err != nil {
		return
	}
	defer func() {
		file.Close()
	}()
}
