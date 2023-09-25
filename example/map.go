package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	m2 := map[string]int{
		"aa": 1,
		"bb": 2,
	}
	// 赋值
	m1["aa"] = 123
	fmt.Printf("%v\n", m1["aa"]) // 123
	fmt.Printf("%v\n", m2["bb"]) // 2

	// map 基础操作
	delete(m2, "bb")
	fmt.Println(m1["bb"]) // 0

	fmt.Println(len(m2)) // 1

	m3 := map[string]int{}
	for i := 0; i < 100; i++ {
		m3[fmt.Sprintf("key-%d", i)] = i
	}
	for key := range m3 {
		fmt.Println(key, "打印")
	}
	for key, value := range m3 {
		fmt.Println(key, value)
	}
	for _, val := range m3 {
		fmt.Println(val, "我是值")
	}
}
