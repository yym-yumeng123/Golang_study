package main

import "fmt"

func main() {
	// 查找 顺序查找

	// 题目1: 有一个数列:
	names := [4]string{"青翼", "紫衫", "金毛", "白眉"}
	var name = ""
	fmt.Println("请输入查找的四大护法")
	fmt.Scanln(&name)

	// 顺序查找 第一种方式
	for i := 0; i < len(names); i++ {
		if name == names[i] {
			fmt.Printf("找到%v, 下标%v", name, i)
			break
		} else if i == (len(names) - 1) {
			fmt.Printf("没有找到%v", name)
		}
	}

	// 顺序查找, 第二种方式
	index := -1
	for i := 0; i < len(names); i++ {
		if name == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到%v, 下标%v", name, index)
	} else {
		fmt.Printf("没有找到%v", name)
	}
}
