package main

import (
	"fmt"
	"sync"
)

func producer(ch chan string, name string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("%s: %d", name, i)
	}
	close(ch)
}

func fanIn(ch1, ch2 chan string) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for {
			select {
			case data, ok := <-ch1:
				if !ok {
					ch1 = nil
					break
				}
				ch <- data
			case data, ok := <-ch2:
				if !ok {
					ch2 = nil
					break
				}
				ch <- data
			}
			if ch1 == nil && ch2 == nil {
				break
			}
		}
	}()
	return ch
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		producer(ch1, "goroutine1")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		producer(ch2, "goroutine2")
	}()

	ch := fanIn(ch1, ch2)

	for data := range ch {
		fmt.Println(data)
	}
}
