package main

import "fmt"

type Usb interface {
	// 声明两个未实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

type Camera struct {
	name string
}

// 让 Phone 实现 usb 接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}
func (p Phone) Call() {
	fmt.Println("手机打电话...")
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
	// 如果 usb 指向 Phone. 调用 call 方法
	// 类型断言
	phone, ok := usb.(Phone)
	fmt.Println(phone, ok)
	if ok == true {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	//c := Computer{}
	//phone := Phone{}
	//camera := Camera{}
	//
	//c.Working(phone)
	//c.Working(camera)

	// 多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"小米"}
	usbArr[1] = Phone{"华为"}
	usbArr[2] = Camera{"sony"}

	var c Computer

	//遍历 usbArr
	for _, v := range usbArr {
		c.Working(v)
	}
	//fmt.Println(usbArr)
}
