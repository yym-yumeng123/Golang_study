package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("C:\\Users\\18026\\Desktop\\Golang_study/example/file/1.txt")
	if err != nil {
		fmt.Println("open file err= ", err)
	}

	fmt.Printf("file=%v", file)

	file.Close()
}
