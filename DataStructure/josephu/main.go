package main

import "fmt"

// 小孩的结构体
type Boy struct {
	No   int
	Next *Boy
}

// 设编号为1, 2, ... n 的n 个人围坐一圈
// 编写一个单向环形链表
// num : 表示小孩的个数,
// *Boy 返回该环形链表的第一个小孩的指针
func AddBoy(num int) *Boy {
	first := &Boy{}  // 空节点
	curBoy := &Boy{} // 空节点 辅助

	if num < 1 {
		fmt.Println("人数不能小于1")
		return first
	}

	for i := 1; i <= num; i++ {
		boy := &Boy{No: i}

		// 分析构成循环链表, 需要一个辅助指针
		if i == 1 { // 第一个小孩
			first = boy // 不要动
			curBoy = boy
			curBoy.Next = first
		} else {
			curBoy.Next = boy
			curBoy = boy
			boy.Next = first
		}
	}

	return first
}

// 显示单向的环形链表
func showBoy(first *Boy) {
	// 处理环形链表为空的情况
	if first.Next == nil {
		fmt.Println("链表为空")
		return
	}

	// 创建一个指针, 帮助遍历 至少有一个小孩
	curBoy := first

	for {
		fmt.Printf("小孩编号=%d -> \n", curBoy.No)

		if curBoy.Next == first {
			break
		}

		curBoy = curBoy.Next
	}
}

// 玩游戏
// 编号为 k (1 < k <= n) 的人从 1开始报数, 数到 m 的那个人出列
// 下一位又从 1开始 报数, 数到 m 的那个人又出列, 一次类推, 直到所有的人出列为止
func playGame(first *Boy, startNo int, countNum int) {
	// 空链表单独出路
	if first.Next == nil {
		fmt.Println("空链表")
		return
	}
	// 定义一个辅助指针
	tail := first
	// 让 tail 指向环形链表的最后一个小孩
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	// 让 first 移动到 startNo [后面删除小孩, 以first为准]
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	// 数到 m 的那个人
	for {
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈\n", first.No)
		// 删除 到 m 的那个小孩
		first = first.Next
		tail.Next = first

		if first == tail {
			break
		}

	}
	fmt.Printf("最后出圈的小孩为 %d", first.No)

}

// 设编号为1, 2, ... n 的n 个人围坐一圈, 约定编号为 k (1 < k <= n) 的人从 1开始报数, 数到 m 的那个人出列,
// 它的下一位又从 1开始 报数, 数到 m 的那个人又出列, 一次类推, 直到所有的人出列为止, 禅城一个出对编号的序列
func main() {
	first := AddBoy(5)

	showBoy(first)

	playGame(first, 2, 3)
}
