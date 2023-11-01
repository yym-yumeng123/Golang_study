package main

import (
	"errors"
	"fmt"
)

func test02() {
	defer func() {
		// recover 内置函数, 可以捕获到异常
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()

	n1 := 10
	n2 := 0
	res := n1 / n2
	fmt.Println("值", res)
}

// 函数去读取配置文后面 init.conf 的信息
// 如果文件名传入不正确, 我们返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		// 读取
		return nil
	} else {
		// 返回一个自定义错误
		return errors.New("读取文件错误")
	}
}

func test03() {
	if err := readConf("config2.ini"); err != nil {
		// 如果读取文件发送错误, 就输出这个错误, 并终止程序
		panic(err)
	}

	fmt.Println("test03继续执行")
}

func main() {
	num1 := 100
	// num1的类型int, num1的值=100, num1的地址0xc00000a0b8
	fmt.Printf("num1的类型%T, num1的值=%v, num1的地址%v\n", num1, num1, &num1)

	num2 := new(int)
	*num2 = 100
	// num2的类型*int, num2的值=0xc00000a0f0, num2的地址0xc00004e028 num2这个指针指向的值0
	fmt.Printf("num2的类型%T, num2的值=%v, num2的地址%v num2这个指针指向的值%v\n", num2, num2, &num2, *num2)

	// 错误处理
	test02()
	fmt.Println("main下面的代码")

	// 自定义错误
	test03()
}
