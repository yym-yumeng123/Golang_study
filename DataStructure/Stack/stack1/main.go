package main

import (
	"errors"
	"fmt"
)

// 使用数组来模拟栈的使用
type Stack struct {
	maxTop int    // 栈最大可以存放个数
	Top    int    // 表示栈顶
	arr    [5]int // 数组模拟栈
}

// 入栈
func (s *Stack) Push(num int) (err error) {
	// 栈是否满了
	if s.Top == s.maxTop-1 {
		fmt.Println("stack full")
		return errors.New("tack full")
	}

	s.Top++
	// 放入数据
	s.arr[s.Top] = num
	return
}

// 遍历栈
func (s *Stack) List() {
	// 先判断栈是否为空
	if s.Top == -1 {
		fmt.Println("stack empty")
	}

	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, s.arr[i])
	}
}

// 出栈
func (s *Stack) Pop() (val int, err error) {
	if s.Top == -1 {
		fmt.Println("栈空了")
		return 0, errors.New("栈空了")
	}
	val = s.arr[s.Top]
	s.Top--
	return val, nil

}

func main() {
	stack := &Stack{
		maxTop: 5,
		Top:    -1, // 当栈顶为 -1, 表示栈空
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.List()
	val, err := stack.Pop()
	fmt.Println(val, err)
	stack.List()
}
