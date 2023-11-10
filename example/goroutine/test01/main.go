package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello, world")
		time.Sleep(time.Second)
	}
}

func main() {

	go test() // 开启一个协程

	for i := 0; i < 10; i++ {
		fmt.Println("Hello, Golang")
		time.Sleep(1000 * time.Millisecond)
	}
}
