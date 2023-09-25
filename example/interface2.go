package main

import "fmt"

type Eater interface {
	Eat()
}

type Animal1 interface {
	Eater
	Sleep()
}

type Dog1 struct {
	Name string
}

func (d Dog1) Eat() {
	fmt.Printf("%s is eating.\n", d.Name)
}

func (d Dog1) Sleep() {
	fmt.Printf("%s is sleep.\n", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Eat() {
	fmt.Printf("%s is eating.\n", c.Name)
}

func (c Cat) Sleep() {
	fmt.Printf("%s is sleep.\n", c.Name)
}

func AnimalEat(a Eater) {
	a.Eat()
}

func main() {
	d := Dog1{Name: "Dog"}
	AnimalEat(d)
}
