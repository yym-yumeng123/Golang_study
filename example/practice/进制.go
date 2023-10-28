package main

import "fmt"

func main() {
	var i int = 5
	// 二进制输出 %b 表示为二进制
	fmt.Printf("%b\n", i)

	// 八进制 0-7 满8进1, 数字0 开头
	var j int = 011
	fmt.Println("j=", j)

	// 0-9 A-F 满16进1, 以0x开头
	var k int = 0x11
	fmt.Println("k=", k)

	// 位运算
	var a int = 1 >> 2
	var b int = -1 >> 2
	var c int = 1 << 2
	var d int = -1 << 2
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)

	fmt.Println("2&3", 2&3)
	fmt.Println("13&7", 13&7)
	fmt.Println("2|3", 2|3)
	fmt.Println("5|4", 5|4)
	fmt.Println("-3^3", -3^3)
}
