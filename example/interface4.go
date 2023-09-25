package main

import "fmt"

// Shape 接口
type Shape interface {
	// area 方法 返回一个 float64类型的值
	area() float64
}

// Rectangle2 结构体
type Rectangle2 struct {
	width  float64
	height float64
}

func (r Rectangle2) area() float64 {
	return r.width * r.height
}

// Circle 结构体
type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var s Shape
	s = Rectangle2{width: 10, height: 20}
	fmt.Printf("矩形面积: %f\n", s.area())

	s = Circle{radius: 3}
	fmt.Printf("圆形面积: %f\n", s.area())
}
