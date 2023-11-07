package service

import "customer/model"

// CustomerService 该结构体完成对 Customer 操作, 包括增删改查
type CustomerService struct {
	customers []model.Customer
	// 当前切片有多少个客户, 可以作为新客户id
	customerNumber int
}

// NewCustomerService 返回 *CustomerService
func NewCustomerService() *CustomerService {
	// 为了能看到有客户在切片中, 初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNumber = 1
	customer := model.NewCustomer(1, "yym", "男", 18, "110", "yym@qq.com")
	customerService.customers = append(customerService.customers, *customer)
	return customerService
}

// List 返回客户切片
func (c *CustomerService) List() []model.Customer {
	return c.customers
}

func (c *CustomerService) AddCustomer(customer model.Customer) bool {
	// 分配id 的规则
	c.customerNumber++
	customer.Id = c.customerNumber
	c.customers = append(c.customers, customer)
	return true
}

// 根据 id 查找用户在切片中的对应下标
func (c *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(c.customers); i++ {
		if c.customers[i].Id == id {
			index = i
		}
	}

	return index
}

// 根据 id 删除客户 (从切片中删除
func (c *CustomerService) Delete(id int) bool {
	index := c.FindById(id)
	if index == -1 {
		return false
	}

	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	return true
}

func main() {

}
