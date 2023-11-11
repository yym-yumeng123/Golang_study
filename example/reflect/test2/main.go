package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	Gender string
}

// 对结构体的反射
func reflectStruct(b interface{}) {
	rType := reflect.TypeOf(b)
	rVal := reflect.ValueOf(b)

	// 获取变量对应的 kind
	fmt.Println("kind", rType.Kind()) // struct

	iv := rVal.Interface()
	fmt.Printf("iv = %v, iv的类型=%T\n", iv, iv) // iv = {张三 12 男}, iv的类型=main.Student

	stu, ok := iv.(Student)
	if ok {
		fmt.Println(stu.Name, stu.Gender)
	}
}

func main() {
	stu := Student{
		Name:   "张三",
		Age:    12,
		Gender: "男",
	}
	reflectStruct(stu)
}
