package main

import (
	"fmt"
	"sync"
	"time"
)

func someFunc() {
	time.Sleep(5 * time.Second)
	fmt.Println("some Func finished")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		someFunc()
	}()
	wg.Wait()
	fmt.Println("main finished")
}
