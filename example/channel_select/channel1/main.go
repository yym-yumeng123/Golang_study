package main

import "fmt"

func main() {
	// 创建一个可以存放 3 个 int 类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan 的值= %v\n", intChan)     // 0xc00006e080 是一个地址
	fmt.Printf("intChan 本身的地址= %v\n", &intChan) // 0xc00004e020

	// 写入数据, 不能超过容量
	intChan <- 10
	num := 211
	intChan <- num

	fmt.Printf("管道的长度len=%v, cap=%v\n", len(intChan), cap(intChan))

	// 读取数据
	var num2 int
	num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("管道的长度len=%v, cap=%v\n", len(intChan), cap(intChan))

	// 在没有使用协程的情况下, 如果管道数据已经全部取出, 再取就会 deadlock
	num3 := <-intChan
	//num4 := <-intChan

	fmt.Println(num3)

}
