package main

import "fmt"

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface
	test03()
}

// 要实现 A, 需要将 B, C的方法都实现
type Stu struct {
	name string
}

func (s Stu) test01() {
	fmt.Println("test01")
}
func (s Stu) test02() {
	fmt.Println("test02")
}
func (s Stu) test03() {
	fmt.Println("test03")
}

type T interface {
}

func main() {
	var s Stu
	var a AInterface = s
	a.test01()
	a.test02()
	a.test03()

	var t T = s
	fmt.Println(t)
}
