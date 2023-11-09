package main

import (
	"fmt"
	"os"
)

func main() {
	file := "C:\\Users\\18026\\Desktop\\Golang_study/example/file/1.txt"
	// 一次性读取文件 ReadFile 把 Open Close 封装到内部了, 不需要管理
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err= %v", err)
	}
	fmt.Printf("%v", string(content))

	fmt.Println("文件读取完成")
}
