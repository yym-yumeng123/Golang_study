package main

import "fmt"

// Phone 接口
type Phone interface {
	// 方法 call
	call()
}

type NokiaPhone struct {
}

func (nokiaphone NokiaPhone) call() {
	fmt.Println("I am Nokia, I call you")
}

type IPhone struct {
}

func (i IPhone) call() {
	fmt.Println("I am iPhone, i can call you")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
