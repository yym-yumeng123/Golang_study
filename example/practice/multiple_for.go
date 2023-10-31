package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Println("i 和 j\n", i, j)
		}
	}

	// 3个班, 每个班 5个同学, 求出各个班的平均分和所有班级的平均分
	for j := 1; j <= 2; j++ {
		sum := 0.0
		for i := 1; i <= 2; i++ {
			var score float64
			fmt.Printf("请输入第%d班 第$%d个学生的成绩 \n", j, i)
			fmt.Scanln(&score)
			// 累积总分
			sum += score
		}
	}

	// 打印金字塔
	var totalLevel int = 9
	for i := 1; i <= totalLevel; i++ {
		// 在打印层数前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}

		// j 表示每层打印多少
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v \t", i, j, i*j)
		}
		fmt.Println("")
	}

	// 随机生成 1-100, 生成了99, 退出
	for {
		num := rand.Intn(100) + 1
		if num == 99 {
			break
		}
		fmt.Println(num)
	}

	// break 指定标签
label1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break // 就近原则 break 默认会跳出最近的 for 循环
				break label1
			}
		}
	}
}
