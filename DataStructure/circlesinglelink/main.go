package main

import "fmt"

type CatNode struct {
	no   int // 编号
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {

	// 判断是否是添加第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head // 构成一个环形
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}

	// 定义一个临时变量. 帮忙找到环形的最后节点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	// 加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

// 输出环形链表
func List(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}

	for {
		fmt.Printf("猫的信息为=[id=%d name=%s]\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

// 删除一个环形单向链表, 如删除 head, 返回一个新 head
func Delete(head *CatNode, id int) *CatNode {
	// 1. 先让 temp 指向 head
	// 2. 让 helper 指向环形链表的最后
	// 3. 让 temp 和 要删除的 id 比较,相同, 同 helper 完成删除

	temp := head
	helper := head // 辅助节点, 指向最后一个节点

	// 空链表的情况
	if temp.next == nil {
		fmt.Println("这是一个空的环形链表, 不能删除")
	}

	// 只有一个节点
	if temp.next == head {
		temp.next = nil
	}

	// helper 指向环形链表的最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	// 两个以上节点
	flag := true
	for {
		// temp 知道head 了, 但temp 本身的 no 还没比较
		if temp.next == head {
			fmt.Println("没有需要删除的节点")
			break
		}
		// 找到了
		if temp.no == id {
			if temp == head { // 说明删除的是头结点
				head = head.next
			}
			helper.next = temp.next
			flag = false
			break
		}

		temp = temp.next     // 不停移动
		helper = helper.next // 移动, 一旦找到要删除的节点 helper
	}
	// 再比较一次
	if flag { // 如果 flag = true, 则我们上面没有删除
		if temp.no == id {
			helper.next = temp.next
		} else {
			fmt.Println("没有这只猫")
		}
	}

	return head
}

func main() {
	// 初始化一个环形链表的头结点
	head := &CatNode{}

	// 创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}

	cat2 := &CatNode{
		no:   2,
		name: "jack",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)

	List(head)
}
