package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 通过反射修改值
	var num int = 10
	reflect1(&num)

	fmt.Println(num)
}

func reflect1(b interface{}) {
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(20)
}
