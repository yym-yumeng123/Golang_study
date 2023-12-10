### 数组

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整型、字符串或者自定义类型

#### 声明数组

```go
/**
arrayName 数组的名称
size 数组的大小
dataType  是数组中元素的数据类型
*/
var arrayName [size]dataType

// 定义了数组 balance 长度为10 类型为 float32
var balance [10]float32
```

#### 初始化数组

在 Go 语言中, 数组的大小是类型的一部分, 因此不同大小的数组是不兼容的, `[5]int` 和 `[10]int` 是不同的类型

初始化数组中 `{}` 中的元素个数不能大于 `[]`中的数字

如果忽略 `[]` 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小

```go
// 对于整数类型, 初始值 0
var numbers [5]int


// 初始化列表来初始化数组的元素
var numbers [5]int{1,2,3,4,5}

// 使用 := 简短声明语法来声明和初始化数组
numbers := [5]int{1,2,3,4,5}


// 定义了数组 balance长度为5类型为 float32,
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.6}

// 字面量
balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.6}

// 如果数组长度不确定, 可以使用 ... 代替数组的长度,
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

// 如果设置了数组的长度, 可以通过指定下标来初始化元素
//  将索引为 1 和 3 的元素初始化
balance := [5]float32{1:2.0, 3:7.0}
```

#### 访问数组元素

数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值

```go
// 读取数组第 10 个元素的值
var salary float32 = balance[9]
```

---

### 指针

变量是一种使用方便的占位符，用于引用计算机内存地址。

Go 语言的取地址符是 `&`, 放到一个变量前使用就会返回相应变量的内存地址

```go
package main
import "fmt"

func main() {
  var a int = 10
  fmt.Println("变量的地址: %x\n", &a) // 变量的地址: 20818a220
}
```

#### 什么是指针

一个指针变量指向了一个值的内存地址

类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下

```go
/**
var-type 为指针类型
var_name 指针变量名
* 用于指定变量作为一个指针
*/
var var_name *var-type

var ip *int // 指向整型
var fp *float32 // 指向浮点型
```

#### 如何使用指针

1. 定义指针变量
2. 为指针变量赋值
3. 访问指针变量中指向地址的值

在指针类型前面加上 \* 号（前缀）来获取指针所指向的内容

```go
package main
import "fmt"

func main() {
  var a int = 20 // 声明实际变量
  var ip *int // 声明指针变量

  ip = &a // 指针变量的存储地址

  fmt.Printf("a 变量的地址是: %x\n", &a  ) // a 变量的地址是: 20818a220
  fmt.Printf("ip 变量储存的指针地址: %x\n", ip ) // ip 变量储存的指针地址: 20818a220
  /* 使用指针访问值 */
  fmt.Printf("*ip 变量的值: %d\n", *ip ) // *ip 变量的值: 20
}
```

#### 空指针

当一个指针被定义后没有分配到任何变量时，它的值为 nil, nil 也称为空指针

nil 在概念上和其它语言的 null、None、nil、NULL 一样，都指代零值或空值

一个指针变量通常缩写为 ptr。

```go
package main

import "fmt"

func main() {
 var ptr *int
 fmt.Printf("ptr 的值为 : %x\n", ptr )
}
```

---

### 结构体

在结构体中我们可以为不同项定义不同的数据类型

结构体是由一系列具有相同类型或不同类型的数据构成的数据集合

#### 定义结构体

结构体定义需要使用 `type 和 struct` 语句。`struct` 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下

```go
type struct_variable_type struct {
  member definition
  member definition
  ...
  member definition
}


// 定义了结构体类型, 就能用于变量的声明
variable_name := structure_variable_type {v1, v2, ...vn}
variable_name := structure_variable_type {key1: v1, key2: v2, ..., keyn: vn}
```

```go
package main
import "fmt"

type Books struct {
  title string
  author string
  subject string
  book_id int
}

func main() {
  // 创建一个新的结构体
  fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

  // 也可以使用 key => value 格式
  fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

  // 忽略的字段为 0 或 空
  fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}
```

#### 访问结构体成员

访问结构体成员, 需要使用 `.` 点号操作符, 格式为:

```go
结构体.成员名

package main
import "fmt"

type Books struct {
  title string
  author string
  subject string
  book_id int
}


func main() {
  var Book1 Books
  var Book2 Books

  Book1.title = "Go语言"
  Book1.author = "www.runoob.com"
  Book1.subject = "Go 语言教程"
  Book1.book_id = 6495407

  /* 打印 Book1 信息 */
  fmt.Printf( "Book 1 title : %s\n", Book1.title)
  fmt.Printf( "Book 1 author : %s\n", Book1.author)
  fmt.Printf( "Book 1 subject : %s\n", Book1.subject)
  fmt.Printf( "Book 1 book_id : %d\n", Book1.book_id)
}

```

#### 结构体作为函数参数

可以像其他数据类型一样将结构体类型作为参数传递给函数。并以以上实例的方式访问结构体变量

```go
package main
import "fmt"

type Books struct {
 title string
 author string
 subject string
 book_id int
}

func main() {
 var Book1 Books    /* 声明 Book1 为 Books 类型 */
 var Book2 Books    /* 声明 Book2 为 Books 类型 */

 /* book 1 描述 */
 Book1.title = "Go 语言"
 Book1.author = "www.runoob.com"
 Book1.subject = "Go 语言教程"
 Book1.book_id = 6495407

 /* book 2 描述 */
 Book2.title = "Python 教程"
 Book2.author = "www.runoob.com"
 Book2.subject = "Python 语言教程"
 Book2.book_id = 6495700

 /* 打印 Book1 信息 */
 printBook(Book1)

 /* 打印 Book2 信息 */
 printBook(Book2)
}

func printBook( book Books ) {
  fmt.Printf( "Book title : %s\n", book.title)
  fmt.Printf( "Book author : %s\n", book.author)
  fmt.Printf( "Book subject : %s\n", book.subject)
  fmt.Printf( "Book book_id : %d\n", book.book_id)
}
```


#### 结构体指针

可以定义指向结构体的指针类似于其他指针变量

```go
var struct_pointer *Books
```

以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前

```go
struct_pointer = &Book1


// 使用结构体指针访问结构体成员, 使用 . 操作符
struct_pointer.title
```

```go
// 结构体自引用, 使用指针
type Node struct {
	left  *Node
	right *Node
}

type Header map[string][]int

var a = string(10) // type(value) 类型转换

// 结构体
func structExample() {
	// sdk1 是 *sdkHttpServer 指针
	sdk1 := &sdkHttpServer{}

	// 实例
	sdk2 := sdkHttpServer{}

	// sdk3 是 *sdkHttpServer
	sdk3 := new(sdkHttpServer)

	// 这样声明, Go 就帮你分配好内存, 不用担心空指针
	var sdk4 sdkHttpServer

	// 就是一个指针
	var sdk5 *sdkHttpServer // nil

	// 赋值, 初始化
	sdk6 := sdkHttpServer{Name: "yym"}
	sdk7 := sdkHttpServer{"yym"}

	println(sdk1, sdk2, sdk3, sdk4, sdk5, sdk6, sdk7)
}

// 指针
func Pointer() {
	// 指针用 * 表示, &表示取地址
	var p *sdkHttpServer = &sdkHttpServer{}
	// 解引用, 得到结构体
	var p1 sdkHttpServer = *p

	// 只是声明, 但没有使用
	var p2 *sdkHttpServer

	fmt.Println(p1, p2)
}

```