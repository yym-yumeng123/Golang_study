package main

import "fmt"

func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}

	// 关闭 intChan
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	// 使用 for
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明该 num 不是素数
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum协程因为取不到数据, 退出")

	// 暂时还不能关闭 primeChan
	exitChan <- true
}

func main() {
	// 统计 1-8000 的数字中, 那些是素数

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) // 放入结果
	// 标识退出的管道
	exitChan := make(chan bool, 4)

	// 开启一个协程, 向 intChan放入 1-8000 个数
	go putNum(intChan)
	// 开启4个协程, 从 intChan 取出数据 并判断是否为素数, 是, 放入到 primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		// 主线程进行处理
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		// 当我们从 exitChan 去取出了 4个结果, 关闭 primeNum
		close(primeChan)
	}()

	// 遍历 primeNum, 把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}

		fmt.Printf("素数=%d\n", res)
	}
}
