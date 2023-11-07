package main

import (
	"customer/model"
	"customer/service"
	"fmt"
)

type customerView struct {
	key  string
	loop bool // 表示是否循环显示菜单
	// 增加一个字段 customerService
	customerService *service.CustomerService
}

func (c *customerView) customerList() {
	// 所有信息都在切片中
	customers := c.customerService.List()
	fmt.Println("------客户列表------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("------客户列表完成------")
}

func (c *customerView) AddCustomer() {
	fmt.Println("------添加客户------")
	fmt.Println("姓名")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("姓别")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomer2(name, gender, age, phone, email)
	c.customerService.AddCustomer(*customer)
	fmt.Println("------添加完成------")
}

func (c *customerView) DeleteCustomer() {
	fmt.Println("------删除客户------")
	fmt.Println("请选择待删除客户编号(-1退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	fmt.Println("确认是否删除(Y/N): ")
	choice := ""
	fmt.Scanln(&choice)

	if c.customerService.Delete(id) {
		fmt.Println("------------删除成功-----------")
	} else {
		fmt.Println("------------删除失败-----------")
	}
}

// 显示主菜单
func (c *customerView) mainMenu() {
	for {
		fmt.Println("------客户信息管理软件------")
		fmt.Println("      1. 添加客户")
		fmt.Println("      2. 修改客户")
		fmt.Println("      3. 删除客户")
		fmt.Println("      4. 客户列表")
		fmt.Println("      5. 退出")
		fmt.Println("请选择(1-5: ")

		fmt.Scanln(&c.key)

		switch c.key {
		case "1":
			c.AddCustomer()
		case "2":
			fmt.Println("修改客户")
		case "3":
			c.DeleteCustomer()
		case "4":
			fmt.Println("客户列表")
			c.customerList()
		case "5":
			c.loop = false
		default:
			fmt.Println("请输入正确的数字")
		}

		if !c.loop {
			break
		}
	}

	fmt.Println("你退出了客户关系系统资料")
}

func main() {
	view := customerView{
		key:  "",
		loop: true,
	}

	view.customerService = service.NewCustomerService()

	view.mainMenu()
}
