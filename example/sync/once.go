package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	data []string
)

func loadData() {
	fmt.Println("loading data") // 只被打印一次
	data = []string{"foo", "bar", "baz"}
}

func getData() []string {
	once.Do(loadData)
	return data
}

func main() {
	fmt.Println(getData())
	fmt.Println(getData())
}
