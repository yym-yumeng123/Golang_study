package main

import (
	"errors"
	"fmt"
)

// 结构体管理环形队列
type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int // 指向队首
	tail    int // 指向队尾
}

func (c *CircleQueue) Push(val int) (err error) {
	if c.IsFull() {
		return errors.New("满了")
	}

	// 分析出 c.tail 在队列尾部, 但是包含最后的元素, 预留一个位置
	c.array[c.tail] = val // 把值给尾部
	c.tail = (c.tail + 1) % c.maxSize
	fmt.Println(c.tail, "c扥")
	return
}

func (c *CircleQueue) Pop() (val int, err error) {
	if c.IsEmpty() {
		return 0, errors.New("队列空了")
	}

	// 取出 head 指向队首, 并且含队首元素
	val = c.array[c.head]
	c.head = (c.head + 1) % c.maxSize
	return
}

// 显示队列
func (c *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下: ")
	// 取出当前队列有多少个元素
	size := c.Size()
	if size == 0 {
		fmt.Println("队列空了")
	}

	// 设计一个辅助的变量, 指向 head
	tempHead := c.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, c.array[tempHead])
		tempHead = (tempHead + 1) % c.maxSize
	}
	fmt.Println()

}

// 判断环形队列是否满了
func (c *CircleQueue) IsFull() bool {
	return (c.tail+1)%c.maxSize == c.head
}

// 是否为空
func (c *CircleQueue) IsEmpty() bool {
	return c.tail == c.head
}

// 取出环形队列有多少个元素
func (c *CircleQueue) Size() int {
	return (c.tail + c.maxSize - c.head) % c.maxSize
}

func main() {
	cir := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	cir.Push(1)
	cir.Push(2)
	cir.Push(3)
	cir.Push(4)
	cir.Push(5)
	cir.ListQueue()
	cir.Pop()
	cir.ListQueue()
	cir.Push(6)
	cir.ListQueue()
}
