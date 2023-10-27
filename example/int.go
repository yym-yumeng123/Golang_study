package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 1
	fmt.Println(i)

	var j int8 = -128 // int8范围 // -128 ~ 127
	fmt.Println("j:", j)

	// uint8范围 0-255
	var k uint8 = 255
	fmt.Println("k:", k)

	// int uint rune byte
	var a int = 9000
	var b uint = 1
	var c byte = 255
	var d rune = 12
	fmt.Println("a:", a, "b:", b, "c:", c,
		"d: ", d)

	var e = 100
	fmt.Printf("e的类型 %T", e)
	var f int64 = 10
	// unsafe.Sizeof(f) unsafe包的一个函数, 可以返回 f 变量占用的字节数
	fmt.Printf("f的类型 %T f的字节数是 %d", f, unsafe.Sizeof(f))

	var age byte = 100
	fmt.Println(age, "age")
}
