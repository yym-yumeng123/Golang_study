package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 创建一个 byte类型的 26个元素的数组,分别设置 A-Z, 使用 for 循环访问所有元素并打印出来. 提示: 字符数据运算 A + 1 -> B
	var arrByte [26]byte

	for i := 0; i < 26; i++ {
		arrByte[i] = 'A' + byte(i)
		fmt.Printf("%c", arrByte[i])
	}

	// 请求出一个数组的最大值, 并得到对应的下标
	var intArr [5]int = [...]int{1, -1, 9, 90, 11}
	maxVal := intArr[0]
	maxValIndex := 0

	for i := 1; i < len(intArr); i++ {
		// 然后从第2个元素开始循环比较, 如果发现有更大, 则交换
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal=%v maxValIndex%v", maxVal, maxValIndex)

	// 求出一个数组的和和平均值 for range
	sum := 0
	for _, v := range [5]int{1, 3, 5, 9} {
		sum += v
	}
	fmt.Printf("和%v, 平均值%v\n", sum, float64(sum)/4)

	// 数组翻转 随机生成五个数, 将其翻转打印
	var arr1 [5]int
	for i, _ := range arr1 {
		arr1[i] = rand.Intn(100)
	}

	fmt.Println(arr1)

	temp := 0
	for i, _ := range arr1 {
		temp = arr1[len(arr1)-i-1]
		arr1[len(arr1)-1-i] = arr1[i]
		arr1[i] = temp
	}

	fmt.Println(arr1)
}
