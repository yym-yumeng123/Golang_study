package main

import "fmt"

type Usb interface {
	// 声明两个未实现的方法
	Start()
	Stop()
}

type Phone struct{}

type Camera struct{}

// 让 Phone 实现 usb 接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

// 让 camera 实现 Usb 的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

// 计算机
type Computer struct{}

// 编写一个方法 Working, 接收一个 Usb 接口类型变量
// 所谓实现了Usb接口, 就是指实现了 Usb 接口的所有方法
func (c Computer) Working(usb Usb) {
	// 通过 usb 接口变量来调用 Start Stop 方法
	usb.Start()
	usb.Stop()
}

func main() {
	c := Computer{}
	phone := Phone{}
	camera := Camera{}

	c.Working(phone)
	c.Working(camera)
}
