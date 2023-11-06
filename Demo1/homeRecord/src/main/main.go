package main

import (
	"fmt"
	"home/utils"
)

func main() {
	fmt.Println("这个是面向对象方式完成")
	utils.NewFamilyAccount().MainMenu()
}
