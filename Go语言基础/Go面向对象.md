### struct 和 method

Go 语言中的 `struct` 是一种自定义的数据类型, 用于组织和存储各种不同的类型的数据. 它
类似于其它变成语言的类或对象, 但是没有继承或多态等面向对象的概念

```go
type Person struct {
	Name string
	Age int
	City string
}

p := Person{"Alice", 30, "New York"}
p := Person{
	Age: 30,
	Name: "Alice",
	City: "New York"
}
```

**可见性**

`struct`中成员变量的大小写决定了它们的可见性和访问权限. 如果一个 `struct` 中的成员变量名
以大写字母开头, name它是可导出的(Exported), 也就是说, 它可以被访问. 这种规则也适用于
其它类型的变量 常量 函数和方法