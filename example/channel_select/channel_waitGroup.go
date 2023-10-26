package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{}, 4)
	var wg sync.WaitGroup

	// 有 20个数据, 但是我们希望执行4个数据
	for i := 0; i < 20; i++ {
		i := i
		wg.Add(1)
		// 当第5个往里面写数据, buffer 4 已经满了
		ch <- struct{}{}
		go func() {
			defer func() {
				wg.Done()
				<-ch
			}()
			time.Sleep(2 * time.Second)
			fmt.Printf("goroutinue %d done\n", i)
		}()
	}

	wg.Wait()
	fmt.Printf("all done")
}
