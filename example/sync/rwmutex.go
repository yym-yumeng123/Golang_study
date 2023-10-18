package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int

func main() {
	lock := sync.RWMutex{}

	for i := 0; i < 4; i++ {
		go func() {
			for j := 0; j < 4; j++ {
				lock.RLock()
				fmt.Printf("data: %d\n", counter)
				time.Sleep(1 * time.Second)
				lock.RUnlock()
			}
		}()
	}

	go func() {
		for i := 0; i < 6; i++ {
			lock.Lock()
			fmt.Printf("write data\n")
			counter++
			time.Sleep(4 * time.Second)
			lock.Unlock()
		}
	}()

	time.Sleep(20 * time.Second)
}
