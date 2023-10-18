package main

import (
	"fmt"
	"sync"
)

type MyStruct struct {
	Name string
	Age  int
}

func main() {

	pool := sync.Pool{
		New: func() interface{} {
			return &MyStruct{}
		},
	}

	for i := 0; i < 100; i++ {
		ms := pool.Get().(*MyStruct)
		ms.Name = fmt.Sprintf("Name: %d", i)
		ms.Age = i

		fmt.Println(*ms)
		pool.Put(ms)
	}
}
