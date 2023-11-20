package main

import "fmt"

// 编写一个函数, 找出路
// myMap 同一个地图, 使用引用
// i j 地图的那一个点测试
func SetWay(myMap *[8][7]int, i int, j int) bool {
	// 什么情况, 就找到出路了
	// myMap[6][5] = 2
	if myMap[6][5] == 2 {
		return true
	} else {
		// 继续找
		// 判断墙在哪
		if myMap[i][j] == 0 { // 如果这个点可以探测
			// 假设这个点通, 但需要探测 上下左右
			myMap[i][j] = 2
			if SetWay(myMap, i+1, j) { // i - 1 向下
				return true
			} else if SetWay(myMap, i, j+1) { // 向右
				return true
			} else if SetWay(myMap, i-1, j) { // 向上
				return true
			} else if SetWay(myMap, i, j-1) { // 向左
				return true
			} else { // 上下左右都不通, 思路,
				myMap[i][j] = 3
				return false
			}
		} else {
			// 不能探测 为1, 是墙
			return false
		}
	}
}

func main() {
	// 创建一个二维数组, 模拟迷宫

	// 1. 元素的值为1, 墙
	// 2. 元素为 0 , 还没有走过的一条路, 没有探测的路径
	// 3. 值为 2, 一条通路
	// 4. 值为 3, 曾经走过, 但走不通
	var myMap [8][7]int

	// 先把地图上下设置为墙 1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}

	myMap[3][1] = 1
	myMap[3][2] = 1

	// 先把地图左右设置为墙 1
	for i := 1; i < 7; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}

	fmt.Println(myMap)

	val := SetWay(&myMap, 1, 1)
	fmt.Println(val, "val")

	fmt.Println(myMap, "myMap")
}
