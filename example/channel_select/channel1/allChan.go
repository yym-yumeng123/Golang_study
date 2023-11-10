package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	cat1 := Cat{
		Name: "yym",
		Age:  12,
	}
	cat2 := Cat{
		Name: "张三",
		Age:  12,
	}
	allChan <- cat1
	allChan <- cat2
	allChan <- 10
	allChan <- "yym"

	// 取出
	cat11 := <-allChan
	cat12 := <-allChan
	fmt.Println(cat12, cat11, cat11.(Cat).Name)
}
