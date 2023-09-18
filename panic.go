package main

import "fmt"

type any = interface{}

func someFunc() {
	fmt.Println("someFunc() called")
	// 1. 一旦执行panic, 程序会立即停止当前函数的执行, 向调用者抛出 panic 异常
	panic("AAAAAAAAAAAAAA")

	// 不执行
	fmt.Println("someFunc() finished")
}

func main() {
	// recover 必须在 defer 中调用
	defer func() {
		// 2. panic 往上调用找到了 recover,返回该异常的值, 没有 panic 异常, 返回 nil
		s := recover()
		println("catch panic: ", s.(string))
	}()

	someFunc()
	fmt.Println("main finished")
}
