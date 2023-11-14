package main

import (
	"errors"
	"fmt"
)

// 数据存入队列时称为 addQueue
// 1. 将尾指针往后移, rear + 1, front == rear 空
// 2. 若尾指针 rear 小于等于队列的最大下标 MaxSize - 1. 则将数据存入 rear 所指的数组元素中, 否则无法存入数据, rear == MaxSize - 1

type Queue struct {
	front   int    // 指向队列的首位
	rear    int    // 指向队尾
	array   [5]int // 数组 => 模拟队列
	maxSize int
}

// AddQueue 数据添加到队列
func (q *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if q.rear == q.maxSize-1 {
		return errors.New("队列满了")
	}

	q.rear++ // rear 后移
	q.array[q.rear] = val
	return err
}

// ShowQueue 显示队列, 找到对首, 然后遍历到队尾
func (q *Queue) ShowQueue() {
	// front 不包含对首的元素
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("arr[%d]=%d\n", i, q.array[i])
	}
}

func (q *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if q.front == q.rear {
		return -1, errors.New("队列空了")
	}

	q.front++ // front 往后移
	val = q.array[q.front]
	return val, err
}

func main() {
	queue := &Queue{
		front:   -1,
		rear:    -1,
		maxSize: 5,
	}

	queue.AddQueue(1)
	queue.AddQueue(2)
	queue.AddQueue(3)
	queue.AddQueue(4)
	queue.AddQueue(5)
	queue.AddQueue(6)
	queue.AddQueue(6)

	queue.ShowQueue()

	fmt.Println()
	fmt.Println()
	queue.GetQueue()
	queue.GetQueue()
	queue.GetQueue()
	queue.GetQueue()

	queue.ShowQueue()
}
