package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//os.OpenFile()

	// 创建一个新文件
	filePath := "C:\\Users\\18026\\Desktop\\Golang_study/example/file/write.txt"
	// 打开文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}

	str := "Hello World yym\n"
	// 写入时使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	// 因为writer 带缓存, 因此在调用 WriteString 方法时, 内容先写入缓存, 所以需要调用 Flush 方法, 将缓存的数据真正写入到文件
	writer.Flush()
}
