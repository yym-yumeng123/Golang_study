package model

import "fmt"

type animal struct {
	Name string
	age  int
}

//因为animal结构体首字母小写,因此只能在 model 使用
// 通过工厂模式来解决

func NewAnimal(n string, s int) *animal {
	return &animal{
		Name: n,
		age:  s,
	}
}

// 字段是私有的
func (a *animal) GetAge() int {
	return a.age
}

// 封装
type person struct {
	Name string
	age  int
	sal  float64
}

// 工厂模式封装
func NewPerson(name string) *person {
	return &person{Name: name}
}

// 访问 age sal
func (p person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p person) SetSal(sal float64) {
	if sal >= 3000 && sal < 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确")
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
