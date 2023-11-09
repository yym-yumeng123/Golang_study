package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// Store 序列化
func (m *Monster) Store() bool {
	// 先序列化, 保存到文件
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("marshal err", err)
		return false
	}

	filePath := "E:\\go_test/monster.ser"
	err = os.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// ReStore 反序列化
func (m *Monster) ReStore() bool {
	// 读取文件
	filePath := "E:\\go_test/monster.ser"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// 反序列化
	err = json.Unmarshal([]byte(data), m)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func main() {

}
