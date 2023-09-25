package main

import "fmt"

func mapInts(nums []int, f func(int) int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = f(num)
	}
	return result
}

func filterInts(nums []int, f func(int) bool) []int {
	var result []int

	for _, num := range nums {
		if f(num) {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	val1 := mapInts(nums, func(a int) int {
		return a * 2
	})

	val2 := filterInts(nums, func(a int) bool {
		return a%2 == 1
	})

	fmt.Println(val1)
	fmt.Println(val2)
}
