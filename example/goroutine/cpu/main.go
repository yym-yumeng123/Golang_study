package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 返回本地 cpu 的个数
	cpuNum := runtime.NumCPU()

	fmt.Println(cpuNum)

	runtime.GOMAXPROCS(cpuNum - 1)
}
