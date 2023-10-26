package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	// 有 buffer 是 1, 第2个 goroutine写不进去, 等待数据读出来
	ch := make(chan struct{}, 1)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- struct{}{}
			counter++
			<-ch
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
