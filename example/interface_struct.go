package main

import "fmt"

type Worker interface {
	doWork()
	Start()
}

// BaseWorker 嵌入了 Worker , 实现了 Start 方法
type BaseWorker struct {
	Worker
}

func (b *BaseWorker) Start() {
	fmt.Println("before start")
	b.doWork()
	fmt.Println("finished")
}

// NormalWork 嵌入了 BaseWork, 实现了 doWork 方法, 实现了 Worker interface
type NormalWork struct {
	BaseWorker
}

func (n *NormalWork) doWork() {
	fmt.Println("doWork")
}

func NewNormalWorker() Worker {
	n := &NormalWork{BaseWorker{}}
	n.Worker = n
	return n
}

func main() {
	w := NewNormalWorker()
	w.Start()
}
