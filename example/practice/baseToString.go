package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 基本数据类型 转 string 类型
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string // 空 str
	var str1 string
	var str2 string
	var str3 string

	// 根据 format 参数生成格式化的字符串并返回字符串, %d	表示为十进制
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q", str, str) // str type string str=99

	// %f	有小数部分但无指数部分，如123.456
	str1 = fmt.Sprintf("%f", num2)
	fmt.Printf("str1 type %T str1=%q", str1, str1)

	// %t	单词true或false
	str3 = fmt.Sprintf("%t", b)
	fmt.Printf("str3 type %T str3=%q", str3, str3)

	// %c	该值对应的unicode码值
	str2 = fmt.Sprintf("%c", myChar)
	fmt.Printf("str2 type %T str2=%q\n", str2, str2)

	// string 类型转换为基本数据类型
	var str4 string = "true"
	var b1 bool
	// 函数返回两个值
	b1, _ = strconv.ParseBool(str4)
	fmt.Printf("b1的type %T b1的值%v", b1, b1)
}
