package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 使用数组来模拟栈的使用
type Stack struct {
	maxTop int     // 栈最大可以存放个数
	Top    int     // 表示栈顶
	arr    [20]int // 数组模拟栈
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

// 判断一个字符是不是运算符 [+ - * /], 使用 ASCII 码
func (s *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 运算的方法
func (s *Stack) Cal(n1 int, n2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = n1 * n2
	case 43:
		res = n1 + n2
	case 45:
		res = n2 - n1
	case 47:
		res = n2 / n1
	default:
		fmt.Println("运算度错误")
	}
	return res
}

// 操作符优先级
func (s *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}

	return res
}

func main() {
	// 数栈 存放数, 符号栈 存放操作符
	numberStack := &Stack{
		maxTop: 20,
		Top:    -1,
	}
	operStack := &Stack{
		maxTop: 20,
		Top:    -1,
	}

	exp := "3+2*6-2"
	// 定义一个 index, 帮助扫描 exp
	index := 0

	// 配合运算, 定义需要的变量
	result := 0

	for {
		ch := exp[index : index+1]
		temp := int([]byte(ch)[0])
		if operStack.IsOper(temp) { // 说明是符号
			// 空栈
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					n1, _ := numberStack.Pop()
					n2, _ := numberStack.Pop()
					oper, _ := operStack.Pop()
					result = operStack.Cal(n1, n2, oper)
					// 将计算结果重新入栈
					numberStack.Push(result)
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}

		} else { // 数字
			val, _ := strconv.ParseInt(ch, 10, 64)
			numberStack.Push(int(val))
		}

		// 继续扫描
		// 先判断 index 是否已经扫描到计算表达式最后
		if index+1 == len(exp) {
			break
		}
		index++
	}

	for {
		if operStack.Top == -1 {
			break
		}
		n1, _ := numberStack.Pop()
		n2, _ := numberStack.Pop()
		oper, _ := operStack.Pop()
		result = operStack.Cal(n1, n2, oper)
		// 将计算结果重新入栈
		numberStack.Push(result)
	}

	res, _ := numberStack.Pop()
	fmt.Printf("表达式%s = %v", exp, res)
}
