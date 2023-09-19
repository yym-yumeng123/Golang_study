package main

import "fmt"

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
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}
