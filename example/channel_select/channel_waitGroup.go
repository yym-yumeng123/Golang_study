package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	for i := 0; i < 2; i++ {
		i := i
		go func() {
			defer func() {
				ch <- struct{}{}
			}()
			time.Sleep(2 * time.Second)
			fmt.Printf("goroutinue %d done", i)
		}()
	}

	for i := 0; i < 2; i++ {
		<-ch
	}
	fmt.Printf("all done")
}
