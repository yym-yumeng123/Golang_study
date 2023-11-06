package service

import "customer/model"

// CustomerService 该结构体完成对 Customer 操作
type CustomerService struct {
	customers []model.Customer
	// 当前切片有多少个客户, 可以作为新客户id
	customerNumber int
}

func main() {

}
