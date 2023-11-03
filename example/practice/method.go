package main

import "fmt"

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

type PersonM struct {
	Name string
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// 给Person类型绑定方法
func (p PersonM) test() {
	fmt.Println("test()", p.Name)
}

func (p PersonM) speak() {
	fmt.Println(p.Name, "是一个好人")
}

func (p PersonM) sum() int {
	sum := 0
	for i := 0; i < 1001; i++ {
		sum += i
	}
	return sum
}

func (p PersonM) sum2(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	var p PersonM
	p.Name = "yym"
	p.test()
	p.speak()
	sum := p.sum()
	fmt.Println(sum)
	sum2 := p.sum2(10)
	fmt.Println(sum2)

	var c Circle
	c.radius = 4.0
	fmt.Println("面积是:", c.area())

	var i integer = 10
	i.print()

}
