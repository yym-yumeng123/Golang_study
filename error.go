package main

import (
	"fmt"
)

// 返回一个值和一个错误类型
func add4(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("a < 0 || b < 0, a:%v, b: %v", a, b)
	}
	return a + b, nil
}

func main() {
	result, err := add4(-1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("result:", result)
	}
}
