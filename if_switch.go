package main

import (
	"fmt"
	"io"
	"os"
)

func computeScore() int {
	return 85
}

func main() {
	// if
	var level string
	// score 变量作用域控制在 if 中
	if score := computeScore(); score >= 90 {
		level = "A"
	} else if score >= 80 {
		level = "B"
	} else {
		level = "C"
	}

	println(level)
	file()

	// switch
	finger := 2
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2, 6, 7, 8, 9:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	default:
		fmt.Println("无效数字")
	}
}

func file() {
	file := "main.go"
	if _, err := os.Stat(file); os.IsNotExist(err) {
		// 文件不存在
		println("file not exist")
	} else {
		f, _ := os.Open(file)
		bytes, _ := io.ReadAll(f)
		println(string(bytes))
	}
}
