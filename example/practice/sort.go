package main

import "fmt"

// BubbleSort 冒泡排序
func BubbleSort(arr *[]int) {
	temp := 0
	fmt.Println("排序前", *arr)
	for j := 1; j < len(*arr); j++ {
		for i := 0; i < len(*arr)-j; i++ {
			if (*arr)[i] > (*arr)[i+1] {
				//(*arr)[i+1] = (*arr)[i]
				temp = (*arr)[i]
				(*arr)[i] = (*arr)[i+1]
				(*arr)[i+1] = temp
			}
		}
	}
	fmt.Println("排序后", *arr)

}

func main() {
	// 冒泡排序
	var arr = []int{24, 69, 80, 57, 13, 46, 70}
	BubbleSort(&arr)
}
