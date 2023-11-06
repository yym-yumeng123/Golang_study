package utils

import "fmt"

func main() {

}

type FamilyAccount struct {
	// 声明必要的字段
	// 声明一个变量, 保存用户输入的选项
	key string
	// 声明一个变量,控制是否退出 for
	loop bool
	// 定义一个变量, 记录是否有收支行为
	flag bool
	// 定义账户的余额
	balance float64
	// 每次收支的金额
	money float64
	// 收支说明
	note string
	// 收支详情使用字符串来记录
	details string
}

// 编写工厂模式的构造方法, 返回一个 *FamilyAccount 实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		flag:    true, money: 0.0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说明",
	}
}

// 将显示明细写成一个方法
func (f *FamilyAccount) showDetails() {
	fmt.Println("------当前收支明细记录------")
	if f.flag {
		fmt.Println(f.details)
	} else {
		fmt.Println("当前没有收支, 请来一笔把")
	}
}

// 登记收入
func (f *FamilyAccount) income() {
	fmt.Println("本次收入金额: ")
	fmt.Scanln(&f.money)
	f.balance += f.money
	fmt.Println("本次收入说明: ")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("\n收入\t%v\t\t%v\t%v", f.balance, f.money, f.note)
	fmt.Println(f.details)
	f.flag = true
}

// 支出
func (f *FamilyAccount) pay() {
	fmt.Println("本次支出金额: ")
	fmt.Scanln(&f.money)
	if f.money > f.balance {
		fmt.Println("余额不足")
		return
	}
	f.balance -= f.money
	fmt.Println("本次支出说明")
	fmt.Scanln(&f.note)
	f.details += fmt.Sprintf("\n支出\t%v\t\t%v\t%v", f.balance, f.money, f.note)
	fmt.Println(f.details)
	f.flag = true
}

// 退出
func (f *FamilyAccount) exit() {
	fmt.Println("你确定要退出记账软件吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误, 请重新输入 y/n")
	}
	if choice == "y" {
		f.loop = false
	}
}

// 给结构体绑定响应的方法
func (f FamilyAccount) MainMenu() {
	for {
		fmt.Println("------家庭收支记账软件------")
		fmt.Println("      1. 收支明细")
		fmt.Println("      2. 登记收入")
		fmt.Println("      3. 登记支出")
		fmt.Println("      4. 退出软件")
		fmt.Println("请选择(1-4: ")

		fmt.Scanln(&f.key)

		switch f.key {
		case "1":
			f.showDetails()
		case "2":
			f.income()
		case "3":
			f.pay()
		case "4":
			f.exit()
		default:
			fmt.Println("请输入正确的选项")
		}

		if !f.loop {
			break
		}
	}
}
