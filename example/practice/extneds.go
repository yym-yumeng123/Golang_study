package main

import "fmt"

// 通用字段 Name Age Score
type DStudent struct {
	Name  string
	Age   int
	Score int
}

// Pupil 小学生
type Pupil struct {
	DStudent // 嵌入 DStudent 结构体
}

type Graduate struct {
	DStudent
}

func (d *DStudent) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", d.Name, d.Age, d.Score)
}

func (d *DStudent) SetScore(score int) {
	d.Score = score
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中...")
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试中...")
}

// 大学生 研究生

func main() {
	var pu1 = &Pupil{}
	pu1.DStudent.Name = "yym"
	pu1.DStudent.Age = 18
	pu1.testing()
	pu1.DStudent.SetScore(100)
	pu1.DStudent.ShowInfo()

	var gr = &Graduate{}
	gr.DStudent.Name = "Mary"
	gr.DStudent.Age = 28
	gr.testing()
	gr.DStudent.SetScore(70)
	gr.DStudent.ShowInfo()
}
