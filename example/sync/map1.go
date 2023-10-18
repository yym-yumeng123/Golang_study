package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mm := sync.Map{}

	for i := 0; i < 10; i++ {
		i := i
		go func() {
			// 添加 .Store(key, value)
			mm.Store(i, i)
		}()
	}

	time.Sleep(2 * time.Second)
	// 获取 Load
	v, ok := mm.Load(1)
	if ok {
		fmt.Println(v.(int))
	}

	// 遍历 range
	mm.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return false
	})
}
