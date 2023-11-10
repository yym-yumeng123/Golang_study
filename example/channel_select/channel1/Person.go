package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	var pChan chan *Person
	pChan = make(chan *Person, 2)
	p := &Person{
		Name:    "yym",
		Age:     12,
		Address: "323",
	}

	pChan <- p

	fmt.Println(pChan)
}
