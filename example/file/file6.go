package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file er=%v", err)
	}

	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)

	// 打开 dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	defer dstFile.Close()
	// 通过 dstFile 获取到 writer
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

func main() {
	CopyFile("a", "b")
}
