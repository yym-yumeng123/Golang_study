package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)
	// 声明一个全局互斥锁
	// sync 同步
	// Mutex 互斥
	lock sync.Mutex
)

// test 函数就是计算 n!, 将这个结果放入到 myMap
func test(n int) {
	res := 1
	for i := 0; i <= n; i++ {
		res *= i
	}
	// 加锁
	lock.Lock()
	myMap[n] = res
	// 解锁
	lock.Unlock()
}

func main() {
	// 开启多个协程
	for i := 0; i < 10; i++ {
		go test(i)
	}

	// 主线程休眠 10 s, 是让协程完成任务
	// 如果没有这句话, 主线程很快就退出, m 中还没有结果
	time.Sleep(time.Second * 5)

	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
