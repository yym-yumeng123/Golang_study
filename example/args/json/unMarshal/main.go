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
	// json => struct
	str := "{\"name\":\"yym\",\"age\":18,\"Skill\":\"study\"}\n"
	// 定义一个 monster 实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("反序列化后 monster=%v, monster.name=%v\n", monster, monster.Name)

	// 定义一个 map
	var a map[string]interface{}

	err1 := json.Unmarshal([]byte(str), &a)
	if err1 != nil {
		fmt.Printf("err=%v", err1)
	}
	fmt.Printf("反序列化后 a=%v\n", a)
}
