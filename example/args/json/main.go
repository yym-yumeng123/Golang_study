package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string
}

func main() {
	m := Monster{
		Name:  "yym",
		Age:   18,
		Skill: "study",
	}

	// 将 m 序列化
	data, err := json.Marshal(&m)
	if err != nil {
		fmt.Printf("序列号错误 err=%v", err)
	}
	fmt.Printf("结构体序列化=%v\n", string(data))

	// 定义一个map
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "张三"
	a["age"] = 20

	dataMap, _ := json.Marshal(a)
	fmt.Printf("map序列化后=%v\n", string(dataMap))

	// 切片进行序列化
	var s []map[string]interface{}
	var m1 map[string]interface{}
	var m2 map[string]interface{}
	m1 = make(map[string]interface{})
	m2 = make(map[string]interface{})
	m1["name"] = "小王"
	m1["gender"] = "男"
	m2["name"] = "校长"
	m2["gender"] = "女"
	//s = make([]map[string]interface{}, 2)
	//s[0] = m1
	//s[1] = m2
	s = append(s, m1, m2)

	dataSlice, _ := json.Marshal(s)

	fmt.Printf("slice序列化后=%v\n", string(dataSlice))

	// 基本数据类型序列化
	var n1 float64 = 23.23

	dataF, _ := json.Marshal(n1)
	fmt.Printf("基本序列化后=%v\n", string(dataF))

}
