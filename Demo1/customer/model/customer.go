package model

import "fmt"

// Customer 表示一个客户信息
type Customer struct {
	Id     int
	Name   string
	Age    int
	Gender string
	Phone  string
	Email  string
}

func NewCustomer(id int, name string, gender string, age int,
	phone string, email string) *Customer {
	return &Customer{
		Id:     id,
		Name:   name,
		Age:    age,
		Gender: gender,
		Phone:  phone,
		Email:  email,
	}
}

func NewCustomer2(name string, gender string, age int,
	phone string, email string) *Customer {
	return &Customer{
		Name:   name,
		Age:    age,
		Gender: gender,
		Phone:  phone,
		Email:  email,
	}
}

// GetInfo 返回用户信息
func (c Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", c.Id, c.Name, c.Age, c.Gender,
		c.Phone, c.Email)
	return info
}

func main() {

}
