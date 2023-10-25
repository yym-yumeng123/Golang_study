package main

import (
	"fmt"
	"sync"
)

func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i < 99; i++ {
		ch <- i * i // 往里写
	}
	fmt.Println("goroutine1 done")
}

func main() {
	// make 可以创建 slice  map channel
	// make(map[int]int)
	// 有 buffer 的 channel
	ch := make(chan int, 10)

	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ch, &wg)

	for num := range ch {
		fmt.Printf("%d\n", num)
	}

	wg.Wait()
}
