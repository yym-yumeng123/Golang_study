package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	//ch <- 0
	close(ch)

	ch <- 1

	num, ok := <-ch
	fmt.Println(ok)  // false
	fmt.Println(num) // 0
}
