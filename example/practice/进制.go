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

	// 其他进制转十进制
}
