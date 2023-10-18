package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		// for 循环完成
		go func() {
			fmt.Println(i) // 10 10 10 ...
		}()
	}

	for j := 0; j < 10; j++ {
		go func(iterI int) {
			fmt.Println(iterI, "变量")
		}(j)
	}

	time.Sleep(1 * time.Second)
}
