package main

import "fmt"

type Animal interface {
	Eat()
	Sleep()
}

type Dog struct {
	Name string
}

// Eat Dog指针
func (d *Dog) Eat() {
	fmt.Printf("%s is eating.\n", d.Name)
}

func (d Dog) Sleep() {
	fmt.Printf("%s is sleeping.\n", d.Name)
}

func main() {
	d := &Dog{Name: "Dog"}
	var a Animal = d

	a.Eat()
}
