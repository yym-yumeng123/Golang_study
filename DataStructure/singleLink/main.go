package main

import "fmt"

// 定义 HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode // 指向下一个节点
}

// 向单链表尾部插入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 1. 先找到该链表的最后这个节点
	// 2. 创建一个辅助节点
	temp := head

	for {
		if temp.next == nil {
			break
		}
		temp = temp.next // 让 temp 不断的指向下一个节点
	}

	// 3. 将 newHeroNode 加入到链表的最后
	temp.next = newHeroNode
}

// 根据 no 的编号从小到大插入 有序插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	// 找到适当的节点
	// 创建一个临时节点
	temp := head

	flag := true

	// 让插入的节点的 no,和 tamp 的下一个节点的 no 比较
	for {
		if temp.next == nil { // 到链表的最后
			break
		} else if temp.next.no > newHeroNode.no {
			// new 就应该插入到 temp 后面
			break
		} else if temp.next.no == newHeroNode.no {
			// 说明我们链表中已经有这个 no, 不让插入
			flag = false
			break
		}

		temp = temp.next // 让 temp 不断的指向下一个节点
	}

	if !flag {
		fmt.Println("已经存在 no")
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}
}

func DeleteHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false

	// 找到要删除的节点
	for {
		if temp.next == nil { // 到链表的最后
			break
		} else if temp.next.no == id {
			// 找到了
			flag = true
			break
		}

		temp = temp.next // 让 temp 不断的指向下一个节点
	}

	if flag {
		temp.next = temp.next.next
	}
}

// 显示链表的所有节点信息
func ListLinkNode(head *HeroNode) {
	// 创建一个辅助节点
	temp := head
	// 判断该链表是否是空链表
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}

	// 遍历这个链表
	for {
		fmt.Printf("节点信息[%d, %s, %s] ==> ", temp.next.no, temp.next.name, temp.next.nickname)
		// 判断是否链表尾
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	// 1. 先创建一个头结点
	head := &HeroNode{}

	// 2. 创建一个新的 HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "林冲",
		nickname: "豹子头",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero2)

	ListLinkNode(head)

	// 删除
	fmt.Println()
	DeleteHeroNode(head, 2)
	ListLinkNode(head)
}
