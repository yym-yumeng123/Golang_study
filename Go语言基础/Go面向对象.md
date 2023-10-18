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

### interface

Go语言提供了另外一种数据类型即接口, 它把所有的具有共性的方法定义在一起, 任何其他类型只要实现了这些方法就是实现了这个接口

接口可以让我们将不同的类型绑定到一组公共的方法上, 从而实现多态和灵活的设计

Go语言中的接口是隐式实现的, 也就是说, 如果一个类型实现了一个接口定义的所有方法, 那么它就自动实现了该接口.
因此, 我们可以通过将接口作为参数来实现对不同类型的调用, 从而实现多态

```go
// 定义接口
type interface_name interface {
    method_name1 [return_type]
	method_name2 [return_type]
	...
}

// 定义结构体
type struct_name struct {
	// variables
}

// 实现接口方法
func (struct_name_variable struct_name) method_name1() [return_type] {
  // 方法实现
}
func (struct_name_variable struct_name) method_namen() [return_type] {
  // 方法实现
}
```

指针类型 receiver 使用指向该类型的指针作为接收者, 值类型 `receiver` 使用该类型的值作为接收者. 这两种接收者类型在实现接口时有不同的影响:

1. 使用指针类型接收者实现接口
    如果类型使用指针类型接收者实现接口, 即实现了该接口所定义的方法, 并且这些方法的接收者时该类型的指针, 则只有指向该类型的指针才能被视为实现了该接口的类型
2. 使用值类型接收者实现接口
    如果类型使用值类型接受者实现接口, 则该类型的值和指向该类型的指针都可以被视为实现了该接口的类型

**interface 嵌套与组合**

接口可以嵌套在其它接口中, 从而形成更复杂的接口类型. 称为接口嵌入

```go
type Eater interface {
	Eat()
}

type Animal1 interface {
	Eater
	Sleep()
}
```

**断言和类型转换**

```go
func AnimalEat(a Eater) {
	a.(Animal).Sleep()
}

func main() {
	d := &EaterDog{}
	AnimalEat(d)
}


// 接口类型断言
func AnimalEat(a Eater)  {
	animal, ok := a.(Animal)
	if ok {
		animal.Sleep()
}
}
```


**判断接口是否为nil**


对于一个 interface 来说, 它有两个要素, 一个是 `type T`, 一个是 `value V`, 分别对应 `reflect` 中的 `TypeOf和ValueOf`

```go
func IsNil(i interface{}) {  
   if i == nil {  
      fmt.Println("i is nil")  
      return  
   }  
   fmt.Println("i isn't nil")  
}  

func main() {  
   var sl Animal = nil  
   fmt.Println(reflect.TypeOf(sl), reflect.ValueOf(sl))  
   IsNil(sl)  
}
```





















