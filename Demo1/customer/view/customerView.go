package main

import "fmt"

type customerView struct {
	key  string
	loop bool // 表示是否循环显示菜单
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
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			fmt.Println("客户列表")
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

	view.mainMenu()
}
