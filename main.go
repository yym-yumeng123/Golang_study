package main
import "fmt"

func main()  {
	println("Hello World")


	// 变量
	// var name string = "yym"
	name := "yym_yu"
	println(name)

	const name1 = "张三"
	// 无法修改
	// name1 = "李四"

	// var i = 100
	// var f float32
	// f = float32(i)
	// println(f)

	var i = 100
	var ip *int
	// 通过 & 拿到内存地址
	ip = &i

	println(i)// 100
	// 读取 i 的值
	println(*ip) // 100
	// 修改 i 的值
	*ip = 200
	println(i) // 200
}