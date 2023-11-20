package main

import "fmt"

// 定义 emp
type Emp struct {
	id   int
	name string
	next *Emp
}

// 定义EmpLink 不带表头, 第一个节点就存放雇员
type EmpLink struct {
	Head *Emp
}

// 添加员工的方法, 编号从小到大
func (e *EmpLink) Insert(emp *Emp) {
	cur := e.Head      // 辅助指针
	var pre *Emp = nil // 这是一个辅助指针, 在 cur 前面
	// 如果当前 EmpLink 是一个空链表
	if cur == nil {
		e.Head = emp
		return
	}

	// 如果不是空链表, 给 emp 找到对应位置并插入
	// cur 和 emp比较, 让 pre 保持在 cur 前面
	for {
		if cur != nil {
			if emp.id < cur.id { // 插入的 id < 当前的id
				// 找到了
				break
			}
			pre = cur
			cur = cur.next
		} else {
			break
		}
	}

	// 退出, 看下是在将 emp 添加到链表最后
	pre.next = emp
	emp.next = cur
}

// 显示链表的信息
func (e *EmpLink) ShowLink(no int) {
	if e.Head == nil {
		fmt.Println("链表为空")
		return
	}

	cur := e.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s", no, cur.id, cur.name)
			cur = cur.next
		} else {
			break
		}
	}
	fmt.Println()
}

func (e *EmpLink) FindById(id int) *Emp {
	cur := e.Head
	for {
		if cur != nil && cur.id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.next
	}
	return nil
}

// 定义hashtable,  含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

func (h *HashTable) Insert(emp *Emp) {
	// 使用散列函数, 确定该雇员添加到那个链表
	linkNo := h.HashFun(emp.id)
	// 使用对应的链表添加
	h.LinkArr[linkNo].Insert(emp)
}

func (h *HashTable) ShowAll() {
	for i := 0; i < len(h.LinkArr); i++ {
		h.LinkArr[i].ShowLink(i)
	}
}

func (h *HashTable) FindById(id int) *Emp {
	linkNo := h.HashFun(id) // 确定雇员在那个链表
	return h.LinkArr[linkNo].FindById(id)
}

// 编写一个散列方法
func (h *HashTable) HashFun(id int) int {
	return id % 7 // 得到一个值, 对应链表的下标
}

func main() {
	// 查找雇员信息, 使用哈希表

	var hashTable HashTable
	emp1 := &Emp{
		id:   1,
		name: "张三",
	}

	emp2 := &Emp{
		id:   2,
		name: "李四",
	}

	emp3 := &Emp{
		id:   8,
		name: "李四",
	}

	hashTable.Insert(emp1)
	hashTable.Insert(emp2)
	hashTable.Insert(emp3)

	hashTable.ShowAll()

	e := hashTable.FindById(2)
	fmt.Println(e, "我是")
}
