package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 第一种方案
var counter = atomic.Int64{}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter.Add(1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter.Load())
}
