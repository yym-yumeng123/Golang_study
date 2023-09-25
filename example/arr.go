package main

import "fmt"

func main() {
	var arr [10]int
	var arr1 [10]int = [10]int{1, 2, 3}

	fmt.Println(arr, arr1)
	fmt.Println("%v", len(arr))

	for i := 0; i < len(arr1); i++ {
		println(arr1[i])
	}

	for index, v := range arr1 {
		println(index, "下标")
		println(v, "值")
	}

	add(&arr1)
	fmt.Println(arr1)
}

func add(arr *[10]int)  {
	for i := range arr {
		arr[i] += 1
	}
}