package main

import (
	"errors"
	"fmt"
)

// 返回一个值和一个错误类型
func add4(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("a < 0 || b < 0, a:%v, b: %v", a, b)
	}
	return a + b, nil
}

type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return e.message
}

func doSomething() error {
	return &MyError{"something went wrong"}
}

func main() {
	result, err := add4(-1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("result:", result)
	}

	if err1 := doSomething(); err1 != nil {
		var myErr *MyError
		if errors.As(err1, &myErr) {
			fmt.Println(myErr)
		} else {
			fmt.Println("err is not MyError")
		}
	}

}
