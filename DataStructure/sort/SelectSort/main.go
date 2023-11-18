package main

import "fmt"

func SelectSort(arr *[7]int) {
	// 最大的值和 arr[0] 交换
	// 假设第一次 arr[0] 是最大值, 第二次 arr[1] 是最大值 ....

	for i := 0; i < len(arr); i++ {
		max := arr[i]
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max { // 找到真正的最大值
				max = arr[j]
				maxIndex = j
			}
		}
		// 交换
		if maxIndex != 0 {
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}
	//max := arr[0]
	//maxIndex := 0
	//for i := 1; i < len(arr); i++ {
	//	if arr[i] > max { // 找到真正的最大值
	//		max = arr[i]
	//		maxIndex = i
	//	}
	//}
	//// 交换
	//if maxIndex != 0 {
	//	arr[0], arr[maxIndex] = arr[maxIndex], arr[0]
	//}

	fmt.Println("第一次", arr)
}

func main() {
	// 从大到小排
	arr := [...]int{10, 34, 19, 100, 80, 79, 15}

	SelectSort(&arr)
	fmt.Println(arr)
}
