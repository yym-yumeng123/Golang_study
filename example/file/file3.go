package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开存在的文件
	filePath := "C:\\Users\\18026\\Desktop\\Golang_study/example/file/write.txt"
	// os.APPEND 追加
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	defer file.Close()

	str := "ABC,ENGLISH\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}
