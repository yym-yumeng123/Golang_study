package main

import "fmt"

type Rectangle1 struct {
	Width  int
	Height int
}

type Person1 struct {
	Name string
	Age  int
}

type Employee struct {
	p1     Person1 // 是一个人
	Salary float64 // 工资
}

// Area 面积
func (r Rectangle1) Area() int {
	return r.Width * r.Height
}

// Scale 放大, 声明指针类型
func (r *Rectangle1) Scale(s int) {
	r.Width *= s
	r.Height *= s
}

func PrintPerson(p Person1) {

}

func main() {
	r := Rectangle1{
		Width:  10,
		Height: 20,
	}

	r.Scale(10)

	fmt.Println(r) // 100 200

	e := Employee{
		p1:     Person1{Name: "yym", Age: 19},
		Salary: 1000,
	}
	fmt.Println(e)
	PrintPerson(e.p1)
}
