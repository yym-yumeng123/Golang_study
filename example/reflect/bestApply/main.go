package main

import (
	"fmt"
	"reflect"
)

// 定义了 Monster 结构体
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float64
	Sex   string
}

// 添加 Print 方法
func (m Monster) Print() {
	fmt.Println("start...")
	fmt.Println(m)
	fmt.Println("end...")
}

// 添加 GetSum 方法
func (m Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// Set方法
func (m Monster) Set(name string, age int, score float64, sex string) {
	m.Score = score
	m.Sex = sex
	m.Name = name
	m.Age = age
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	// 对结构体遍历, 判断是否是结构体
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	// 获取结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为%v\n", i, val.Field(i))
		// 获取到 struct 标签, 需要通过 reflect.type 来获取 tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field%d: tag为%v\n", i, tagVal)
		}
	}

	// 获取结构体的方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//var params []reflect.Value
	val.Method(1).Call(nil)

	// 调用结构体的第一个方法 Method[0]
	var params []reflect.Value // 声明一个 reflect 切片
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) // 传入的参数是 []reflect.Value, 返回也是 []reflect.Value
	fmt.Println("res=", res[0].Int())
}

func main() {
	a := Monster{
		Name:  "牛魔王",
		Age:   19,
		Score: 100,
		Sex:   "男",
	}

	TestStruct(a)
}
