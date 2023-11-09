package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	/**
	1. file 叫 file对象
	2. file 叫 file指针
	3. file 叫 file文件句柄
	*/
	file, err := os.Open("C:\\Users\\18026\\Desktop\\Golang_study/example/file/1.txt")
	// 及时关闭 file 句柄, 否则会有内存泄露
	defer file.Close()
	if err != nil {
		fmt.Println("open file err= ", err)
	}

	// 创建一个 *Reader, 是带缓冲的 默认 defaultBufSize = 4096,
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {                  // io.EOF表示文件的末尾
			break
		}
		fmt.Print(str)
	}

	fmt.Println("文件读取完成")
}
