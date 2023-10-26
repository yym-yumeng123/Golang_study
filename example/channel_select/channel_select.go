package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}()

	ticker := time.NewTicker(100 * time.Millisecond)

loop:
	for {
		select {
		case num, ok := <-ch:
			if !ok {
				break loop
			}
			fmt.Printf("receive num: %d\n", num)
		case <-ticker.C:
			fmt.Println("默认")

			//default:
		}
		//time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
}
