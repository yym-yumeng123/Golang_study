package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mm := make(map[int]int)
	var rwLock sync.RWMutex

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			rwLock.Lock()
			defer rwLock.Unlock()
			mm[i] = i
		}()
	}

	time.Sleep(2 * time.Second)
	for k, v := range mm {
		fmt.Println(k, v)
	}
}
