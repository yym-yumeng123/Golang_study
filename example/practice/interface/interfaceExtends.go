package main

import "fmt"

type BirdAble interface {
	flying()
}

type Fish interface {
	swimming()
}

type Monkey struct {
	Name string
}

// LittleMonkey 小猴子
type LittleMonkey struct {
	Monkey
}

func (m *Monkey) climbing() {
	fmt.Println("生来会爬树", m.Name)
}

func (m *LittleMonkey) flying() {
	fmt.Println("会飞了", m.Name)
}

func (m *LittleMonkey) swimming() {
	fmt.Println("会游泳了", m.Name)
}

func main() {
	little := &LittleMonkey{
		Monkey{
			Name: "小红",
		},
	}

	little.climbing()
	little.flying()
	little.swimming()

	var x interface{}
	var b float64
	x = b
	b = x.(float64)
}
