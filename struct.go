package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	City string
}

type Rectangle struct {
	Width  int
	Height int
}

func main() {
	p := Person{
		Name: "yym",
		Age:  18,
		City: "bj",
	}
	fmt.Println(p.Name)

	r := Rectangle{Width: 10, Height: 10}
	fmt.Println(r.Area())

	p1 := Person{
		Name: "yym1",
		Age:  19,
		City: "SH",
	}
	p2 := p1 // 复制
	p1.Name = "yym1_修改"
	fmt.Println(p1, p2)

	p3 := &Person{
		Name: "yym3",
		Age:  21,
		City: "SZ",
	}
	p4 := p3
	p3.Name = "yym3_修改"
	fmt.Println(p3, *p4)
}

// Area struct method
func (r Rectangle) Area() int {
	return r.Width * r.Height
}
