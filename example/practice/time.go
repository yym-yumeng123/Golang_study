package main

import (
	"fmt"
	"strconv"
	"time"
)

func test01() {
	str := ""

	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now) // now=2023-11-01 17:58:34.9529612 +0800 CST m=+0.004471701 now type=time.Time

	// 通过now获取年月日时分秒
	fmt.Printf("当前时间是%d年%d月%d日 %d:%d\n", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute())

	// 格式化时间
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println()

	const (
		Nanosecond  = 1                  // 纳秒
		Microsecond = 1000 * Nanosecond  // 微妙
		Millisecond = 1000 * Microsecond // 毫秒
	)

	// 每隔 1 s 打印一个数字, 打印到 100 就退出
	i := 0
	for {
		i++
		fmt.Println(i)
		// 休眠
		time.Sleep(10 * time.Millisecond)
		if i == 100 {
			break
		}
	}

	// unix UnixNano 1970
	// unix时间戳=1698834112, unixNano时间戳=1698834112400299100
	fmt.Printf("unix时间戳=%v, unixNano时间戳=%v\n", now.Unix(), now.UnixNano())

	// 函数执行的时间
	start := time.Now().Unix()
	test01()
	end := time.Now().Unix()
	fmt.Printf("执行事件%v秒", end-start)
}
