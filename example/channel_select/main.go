package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// make 可以创建 slice  map channel
	// make(map[int]int)
	// 无 buffer 的 channel
	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		time.Sleep(5 * time.Second) // 阻塞
		defer wg.Done()
		ch <- 1 // 往里写
		fmt.Println("goroutine1 done")
	}()

	num := <-ch
	fmt.Printf("%d\n", num)

	wg.Wait()
}
