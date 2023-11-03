package main

import (
	"encoding/json"
	"fmt"
)

// Cat 结构体
type Cat struct {
	Name  string
	Age   int
	Color string
}

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int              // 指针 初始值: nil
	slice  []int             // 切片 // 初始 nil []  需要先 make 才能用
	map1   map[string]string // map 初始 nil 也需要先 make
}

type Monster struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 创建一个Cat的变量
	var cat1 Cat = Cat{Name: "yym", Age: 12, Color: "yellow"}
	fmt.Println(cat1)

	var p1 Person
	fmt.Println(p1)

	var mon1 Monster = Monster{Name: "牛魔王", Age: 12}
	mon2 := mon1
	mon2.Name = "青牛" // 结构体是值类型, 默认值拷贝
	fmt.Println("m1", mon1)
	fmt.Println("m2", mon2)

	// 创建方式3
	var m1 *Monster = new(Monster)
	// 因为 m1 是指针, 因此标准的给字段赋值
	(*m1).Name = "yym"
	(*m1).Age = 15
	// 上面写法等价于
	m1.Name = "John"

	fmt.Println("m1", *m1)

	// 创建方式4 var m2 *Mon = &Mon{}
	var m2 *Monster = &Monster{
		Name: "yym1",
		Age:  10,
	}
	(*m2).Name = "yym1_修改"
	fmt.Println("m2", *m2)

	// 将 struct 变量进行 json处理 问题: json处理后字段名也是首字母大写, 返回给前端, 不习惯
	// Monster 序列化 json格式 -> 字符串
	var monster Monster = Monster{"牛魔王", 500}
	// 序列化成json
	jsonStr, _ := json.Marshal(monster)
	fmt.Println("jsonStr", string(jsonStr)) //  {"Name":"牛魔王","Age":500}

}
