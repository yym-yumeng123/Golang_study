package model

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

func main() {

}
