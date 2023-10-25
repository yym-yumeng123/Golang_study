package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	close(ch)

	for num := range ch {
		fmt.Printf("num: %d\n", num)
	}

	var ch1 chan int
	ch1 <- 20
}
