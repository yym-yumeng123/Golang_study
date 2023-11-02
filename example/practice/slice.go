package main

import "fmt"

func main() {
	var intArr [5]int = [...]int{1, 2, 3, 4, 5}

	// 声明定义一个切片
	/**
	 * slice 切片名, intArr[1:3] 表示 slice 引用到 intArr 这个数组
	 * 左闭右开原则, 包括左边不包右边 下标1开始到3
	 */
	slice := intArr[1:3] // 2 3
	fmt.Println("intArr", intArr)
	fmt.Println("slice元素", slice)
	fmt.Println("slice元素个数", len(slice))
	fmt.Println("slice元素的容量", cap(slice)) // 容量动态变化

	// 使用 make 创建切片
	/**
	 * 可以指定切片的大小和容量
	 * 没有给切片的各个元素赋值, 会使用默认值
	 * 由 make 方式创建的切片对应数组由make底层维护,对外不可见, 只能通过 slice 访问
	 */
	var slice1 []float64 = make([]float64, 5, 10)
	slice1[0] = 10
	slice1[2] = 20
	fmt.Println(slice1)
	fmt.Println("slice1的size", len(slice1))

	// 直接指定具体数组
	var slice2 []string = []string{"tom", "jack", "mary"}
	fmt.Println(slice2)
	fmt.Println("slice2的size", len(slice2))

	// for
	for i := 0; i < len(slice2); i++ {
		fmt.Println("我是", slice2[i])
	}

	// for range
	for _, v := range slice2 {
		fmt.Printf("v=%v\n", v)
	}

	// append 内置函数
	var slice3 []int = []int{100, 200, 300}
	fmt.Println("slice3", slice3)
	slice3 = append(slice3, 400)
	slice3 = append(slice3, 500)
	fmt.Println("slice3", slice3)
	// 通过append将切片追加给自己, 第二个参数必须是切片
	slice3 = append(slice3, slice3...)
	fmt.Println("slice3", slice3)

	// copy 内置函数
	var a []int = []int{1, 2, 3, 4, 5}
	var slice4 = make([]int, 10)
	copy(slice4, a)
	fmt.Println("a", a)
	fmt.Println("slice4", slice4)

	// string slice
	var str = "hello@guigu"

	// 使用切片
	slice5 := str[6:]
	fmt.Println("slice5", slice5)
	//str[0] = 'z'
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)
	str1 := "中国"
	arr2 := []rune(str1)
	arr2[0] = '爱'
	str1 = string(arr2)
	fmt.Println("str1=", str1)

}
