package main

import "fmt"

func IsNil(i interface{}) {
	if i == nil {
		fmt.Println("i is nil")
		return
	}
	fmt.Println("i isn't nil")
}

func main() {
	IsNil(nil)
}
