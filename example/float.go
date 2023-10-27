package main

import "fmt"

func main() {
	var price float32 = 89.32
	fmt.Println(price)

	var num1 float32 = -0.00089
	var num2 = -743434.09
	fmt.Println("num1:", num1, num2)
	fmt.Printf("%T", num2)

	f1 := 12
	f2 := .12
	fmt.Println(f1, f2)

	// 字符char类型 byte -> uint8 0-255
	// 如果我们保存的字符在 ASCII 表的, 比如 [0-1a-zA-Z..] 直接保存 byte
	var char1 byte = 'a'
	var char2 byte = '0'
	var char3 int = '北'
	var c4 byte = '\n'
	fmt.Println("char1=", char1) // 输出的是对应字符的ASCII码值 char1= 97
	fmt.Println("char2=", char2) // char1= 48
	fmt.Printf("c1= %c, c3= %c", char1, char3)
	fmt.Println("c4的\n", c4)

	var n1 = 10 + '中'
	fmt.Printf("n1=%c\n", n1)
	fmt.Println("n1=", n1)
}
