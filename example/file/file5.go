package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开存在的文件
	filePath := "C:\\Users\\18026\\Desktop\\Golang_study/example/file/write.txt"
	// os.TRUNC 读 添加
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	defer file.Close()

	// 先读取
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	str := "你好我是yym\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
