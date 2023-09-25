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

p.Name // 使用
p.Name = "yym" // 修改
```

**可见性**

`struct`中成员变量的大小写决定了它们的可见性和访问权限. 如果一个 `struct` 中的成员变量名
以大写字母开头, name它是可导出的(Exported), 也就是说, 它可以被访问. 这种规则也适用于
其它类型的变量 常量 函数和方法

### struct 和 method

结构体(struct)的方法可以定义为值接收器或指针接收器, 这两种接收器有一些区别

**接收器**

在 Go 语言中, 方法是一种与类型相关联的函数. 方法的定义包括方法名, 接收器类型和函数体.
接收器类型指定了该方法所属的类型,并且接收器是值类型或指针类型. 在方法内部,可以使用接收器访问和
修改类型的值

接收器的原理与函数的参数是相同的, 他们都是用来接受函数调用时传递的参数, 不同之处在于,
接收器作为方法的参数, 必须放在函数名之前, 并且使用特殊的语法来指定接收器的类型

```go
type Person struct {
	Name string
	Age int
}

func (p Person) Print()  {
    fmt.Println(p.Name, p.Age)
}

func main() {
    p := Person{"aa", 18}
		Person.Print(p)
}
```

**值接收器**

值接收器使用实例的副本来调用方法. 在方法内部, 对实例的修改不会影响原始实例. 值接收器适用于对实例进行只读操作的方法

```go
type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

rect := Rectangle{Width:180, Height: 20}
fmt.Println(rect.Area())
```

**指针接收器**

指针接收器使用实例的指针来调用方法。在方法内部，对实例的修改会影响原始实例。指针接收器适用于对实例进行读写操作的方法。

```go
type Rectangle struct {
	Width  int
	Height int
}

func (r *Rectangle) Scale(factor int)  {
	r.Width *= factor
	r.Height *= factor
}
rect := &Rectangle{Width: 10, Height: 20}
rect.Scale(2)
fmt.Println(rect.Width, rect.Height) // Output: 20 40
```

**传递结构体值作为函数参数**

如果将结构体的值传递给函数，那么函数将复制整个结构体的值，并在栈上创建一个新的结构体实例。这种方式会导致内存的开销增加，特别是当结构体实例很大时，内存的开销会很高。

```go
type Rectangle struct {
    Width  int
    Height int
}

func (r Rectangle) Print() {
    fmt.Printf("width=%d, height=%d\n", r.Width, r.Height)
}

rect := Rectangle{Width: 10, Height: 20}
rect.Print()
```

**传递结构体指针作为函数参数**

如果将结构体的指针传递给函数，那么函数将复制指针，并在栈上创建一个新的指针变量。这种方式不会复制整个结构体，只会复制一个指向结构体的指针，因此内存的开销会很小。

```go
type Rectangle struct {
    Width  int
    Height int
}

func (r *Rectangle) Scale(factor int) {
    r.Width *= factor
    r.Height *= factor
}

rect := &Rectangle{Width: 10, Height: 20}
Scale(rect, 2)

func Scale(rect *Rectangle, factor int) {
    rect.Width *= factor
    rect.Height *= factor
}
```

### 结构体的嵌套

- 如果嵌入的两个结构体类型中有相同的字段名, 则在访问这个字段时会出现命名冲突. 为了避免, 可以通过给字段名加上结构体类型前缀的方式来解决
- 如果嵌入的结构体类型有私有字段(字段名以小写字母开头), 则字段在外部结构体不可访问, 如果需要访问, 可以在外部结构体中定义一个公开的方法来访问它们
- 注意嵌入结构体类型是指针还是值
- 作为函数参数时, 外层结构体不能作为内层结构体类型使用

```go
type Person struct {
	Name string
	Age int
}


type Employee struct {
	Person      // 是一个人
	Salary  float64 // 工资
}
```