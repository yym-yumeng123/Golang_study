package main

import "fmt"

func lessThan100(i int) bool {
	return i < 100
}

func main() {
	for i := 0; lessThan100(i); i++ {
		fmt.Println(i)
	}

	for i := 0; lessThan100(i); {
		fmt.Println(i)
		i++
	}

	i := 0
	for lessThan100(i) {
		fmt.Println(i)
		i++
	}

	for {
		if !lessThan100(i) {
			break
		}
		fmt.Println(i)
		i++
	}

	// break continue
	for i := 0; i < 10; i++ {
		if i == 4 {
			break // 中止后面的操作
		}
		fmt.Println(i) // 0 1 2 3
	}

	for i := 0; i < 10; i++ {
		if i == 4 {
			continue // 中止当前操作
		}
		fmt.Println(i) // 0 1 2 3 5 6 7 8 9
	}

	// 多重循环
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer
			}
			fmt.Printf("(%d * %d)", i, j)
		}
	}

	number := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for idx, value := range number {
		fmt.Println(idx, value)
	}

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 变量问题
	numbers := []*int{}
	var oddNumbers []*int
	for i := 1; i < 8; i++ {
		num := i // 创建新的变量
		numbers = append(numbers, &num)
	}
	for _, number := range numbers {
		if (*number)%2 == 1 {
			oddNumbers = append(oddNumbers, number)
		}
	}

	for _, oddNumber := range oddNumbers {
		fmt.Println(*oddNumber)
	}
}
