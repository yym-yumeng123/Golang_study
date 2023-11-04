package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1. 声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2.声明一个Hero结构体切片类型
type HeroSlice []Hero

// 3. 实现 Interface 接口
func (h HeroSlice) Len() int {
	return len(h)
}

// 按Hero的年龄从小到大排列
func (h HeroSlice) Less(i, j int) bool {
	return h[i].Age < h[j].Age
}

func (h HeroSlice) Swap(i, j int) {
	temp := h[i]
	h[i] = h[j]
	h[j] = temp
}

func main() {
	// 实现对 Hero 结构体切片的排序: sort.Sort(data interface)

	// 定义一个切片
	var intSlice = []int{0, -1, 10, 7, 90}
	// 使用系统提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	// 结构体切片排序
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	for _, i2 := range heros {
		fmt.Println("i2", i2)
	}

	sort.Sort(heros)

	fmt.Println(heros)
}
