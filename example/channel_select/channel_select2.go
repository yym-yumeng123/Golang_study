package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// 不能在 `select` 语句中使用相同的 `unbuffer channel` 进行读写操作
	select {
	case x := <-ch:
		fmt.Printf("num: %d\n", x)
	case ch <- 0:
	case <-time.After(1 * time.Second):
		break
	}
}
