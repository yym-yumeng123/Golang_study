package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ready := false
	var lock sync.Mutex
	cond := sync.NewCond(&lock)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Printf("goroutine %d start\n", i)
			for {
				lock.Lock()
				if ready {
					lock.Unlock()
					break
				}
				cond.Wait()
				lock.Unlock()
			}
			fmt.Printf("goroutine %d done\n", i)
		}()
	}

	time.Sleep(2 * time.Second)
	ready = true
	//cond.Signal() // 唤醒一个 goroutine
	cond.Broadcast()
	time.Sleep(5 * time.Second)
}
