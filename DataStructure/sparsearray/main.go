package main

import "fmt"

type ValNode struct {
	row int // 列
	col int // 行
	val int // 值
}

func main() {
	// 1. 创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 // 白字
	// 2. 原始数组
	for _, v := range chessMap {
		fmt.Println(v)
	}

	// 3. 转成稀疏数组
	// 遍历, 有一个元素的值不为 0, 创建一个 node 结构体,
	// 放入对应的切片
	var sparseArr []ValNode

	// 标准的稀疏数组含有一个表示记录原始二维数组的规模(行数, 列数, 默认值)
	firstValNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}

	sparseArr = append(sparseArr, firstValNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				sparseArr = append(sparseArr, ValNode{
					row: i,
					col: j,
					val: v2,
				})
			}
		}
	}

	fmt.Println(sparseArr)

	// 创建一个原始数组
	var arr [11][11]int
	// 4. 将稀疏数组 sparseArr 恢复
	for i, valNode := range sparseArr {
		if i != 0 {
			arr[valNode.row][valNode.col] = valNode.val
		}
	}

	for _, v := range arr {

		fmt.Println(v, "arr")
	}
}
