package main

import (
	"fmt"
)

func WriteData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
	}
	close(intChan) // 关闭
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData读到数据=%v\n", v)
	}

	exitChan <- true
	close(exitChan)
}

func main() {
	// 创建两个管道
	var intChan chan int = make(chan int, 50)
	var exitChan chan bool = make(chan bool, 1)

	go WriteData(intChan)
	go readData(intChan, exitChan)
	//time.Sleep(time.Second * 5)
	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
