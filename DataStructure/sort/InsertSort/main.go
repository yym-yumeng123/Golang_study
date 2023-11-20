package main

import "fmt"

func InsertSort(arr *[5]int) {

	for i := 1; i < len(arr); i++ {
		// 完成第 i 次, 给第 i + 1个元素找到合适的位置并插入
		insertVal := arr[i]  // 第 i 次无序的元素
		insertIndex := i - 1 // 下标

		// 从大到小 插入的值 > 有序表
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}

		// 插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}

	fmt.Println("第一次", *arr)

	//// 完成第二次, 给第三个元素找到合适的位置并插入
	//insertVal = arr[2]  // 第二次第一个无序的元素
	//insertIndex = 2 - 1 // 下标
	//
	//// 从大到小 插入的值 > 要插入的值
	//for insertIndex >= 0 && arr[insertIndex] < insertVal {
	//	arr[insertIndex+1] = arr[insertIndex] // 数据后移
	//	insertIndex--
	//}
	//
	//// 插入
	//if insertIndex+1 != 2 {
	//	arr[insertIndex+1] = insertVal
	//}
	//
	//fmt.Println("第一次", *arr)

}

func main() {
	// 插入排序
	arr := [...]int{23, 0, 12, 56, 34}

	InsertSort(&arr)

}
