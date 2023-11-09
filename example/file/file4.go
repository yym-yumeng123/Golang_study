package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开存在的文件
	filePath := "C:\\Users\\18026\\Desktop\\Golang_study/example/file/write.txt"
	// os.TRUNC 清空
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	defer file.Close()

	str := "你好我是yym\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}
