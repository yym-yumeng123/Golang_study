package main

import "fmt"

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}

	middleIndex := (leftIndex + rightIndex) / 2

	if (*arr)[middleIndex] > findVal {
		// 范围应该在 leftIndex => middleIndex - 1
		BinaryFind(arr, leftIndex, middleIndex-1, findVal)
	} else if (*arr)[middleIndex] < findVal {
		BinaryFind(arr, middleIndex+1, rightIndex, findVal)
	} else {
		fmt.Println("我要找到是", findVal, middleIndex)
	}

}

func main() {
	// 二分查找
	arr := [6]int{1, 8, 10, 89, 100, 123}
	BinaryFind(&arr, 0, len(arr)-1, 8)
}
