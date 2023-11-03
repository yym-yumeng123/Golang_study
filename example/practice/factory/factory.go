package main

import (
	"fmt"
)

//type animal struct {
//	Name string
//	Age  int
//}

//var NewAnimal = NewAnimal

func main() {
	//var s = model.Animal{
	//	Name: "yym",
	//	Age:  18,
	//}

	var a = NewAnimal("yym", 18)
	fmt.Println(a)
}

// 应该引入 model/NewAnimal包的, mod 后面搞
func NewAnimal(s string, i int) interface{} {
	return "random"
}
