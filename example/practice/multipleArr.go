package main

import "fmt"

func main() {
	// 二维数组
	/**
	000000
	001000
	020300
	000000
	*/

	// 定义/声明二维数组
	var arr [4][6]int
	// 赋值
	arr[2][3] = 1
	arr[3][1] = 2
	arr[3][4] = 3

	// 遍历二维数组
	for _, v := range arr {
		for _, subV := range v {
			fmt.Print(subV, " ")
		}
		fmt.Println()
	}

	fmt.Println(arr)

	var arr3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var arr4 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var arr5 = [...][2]int{{1}, {2}, {3}}

	fmt.Println(arr5, arr4, arr3)

	// 定义二维数组是,用于保存 三个班, 每个班五名同学的成绩
	var scores = [3][5]float64{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 14, 15, 16}}

	fmt.Println(scores, "scores")
}
