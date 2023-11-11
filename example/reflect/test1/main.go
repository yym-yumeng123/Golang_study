package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {
	// 通过反射获取传入的变量的 type kind
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	// 获取 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

	iv := rVal.Interface()
	num2 := iv.(int)

	fmt.Println(num2)

}

func main() {
	// 对 基本数据类型 interface{} reflect.value 的基本操作
	var num int = 100
	reflectTest(num)
}
