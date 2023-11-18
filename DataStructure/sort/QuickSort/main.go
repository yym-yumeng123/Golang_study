package main

import "fmt"

/*
*
1. left 数组左边的下标
2. right 数组右边的下标
3. arrat 表示要排序的数组
*/
func QuickSort(left int, right int, array *[6]int) {
	l := left
	r := right
	// pivot 中轴
	pivot := array[(left+right)/2]
	temp := 0

	for l < r {
		// 先从中间的左边 找到比 pivot 大的值, 只要小于, 就一直找下标的下一个值
		for array[l] < pivot {
			l++
		}
		// 从中间的右边 找到 小于 pivot 的值,
		for array[r] > pivot {
			r--
		}
		// l >= r, 表明本次分解任务完成
		if l >= r {
			break
		}

		// 交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp

		//优化
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}

	if l == r {
		l++
		r--
	}

	// 向左递归
	if left < r {
		QuickSort(left, r, array)
	}

	if right > l {
		QuickSort(l, right, array)
	}

}

func main() {
	// 快速排序 从小到大
	arr := [...]int{-9, 78, 0, 23, -567, 70}

	// 调用
	QuickSort(0, len(arr)-1, &arr)

	fmt.Println(arr)
}
