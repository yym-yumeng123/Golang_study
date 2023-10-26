package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch1)
		for i := 0; i < 10; i++ {
			ch1 <- i
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch2)
		for i := 0; i < 10; i++ {
			ch2 <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	timeout := time.After(7 * time.Second)
loop:
	for {
		select {
		case num, ok := <-ch1:
			if !ok {
				ch1 = nil
				break
			}
			fmt.Printf("receive data from ch1: %d\n", num)
		case num, ok := <-ch2:
			if !ok {
				ch2 = nil
				break
			}
			fmt.Printf("receive data from ch2: %d\n", num)
		case <-timeout:
			fmt.Println("no data, timeout")
			break loop
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}

	wg.Wait()
}
