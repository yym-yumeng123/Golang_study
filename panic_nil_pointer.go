package main

import "fmt"

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	// 空指针
	var ip *int

	// 直接使用空指针 invalid memory address or nil pointer dereference
	fmt.Println(*ip) // recover 捕获错误程序可以正常执行

	var s []int = []int{1, 2, 3}
	fmt.Println(s[10], "s")
}
