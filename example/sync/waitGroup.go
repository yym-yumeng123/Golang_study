package main

import (
	"fmt"
	"sync"
)

func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	// 模拟一些工作
	for i := 0; i < 5; i++ {
		fmt.Printf("Worker %d working\n", id)
	}

	fmt.Printf("Worker %d done\n", id)
}
func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go Worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")
}
