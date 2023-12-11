### 创建 "Hello World"

```go
// main.go 全小写字母形式命名, .go 扩展名为结尾

// 定义包名 
// 整个 Go 程序中仅允许存在一个名为 main 的包
package main

// 导入包
import "fmt"

// main 函数
// 当你运行一个可执行的 Go 程序的时候，所有的代码都会从这个入口函数开始运行
func main() {
	// fmt 输出到终端
	fmt.Println("Hello World")
}
```

```shell
# 运行 main.go
go run main.go

# 编译 main.go, 获得一个二进制可执行文件
go build main.go
./main.go
```

### Go Module

Go 语言的构建模式历经了三个迭代和演化过程，分别是最初期的 `GOPATH`、1.5 版本的 `Vendor` 机制，以及现在的 `Go Module`

`Go module` 构建模式是在 Go 1.11 版本正式引入的

Go 1.16 版本中，Go module 已经成为了 Go 默认的包依赖管理机制和 Go 源码构建机制

Go Module 的核心是一个名为 go.mod 的文件，在这个文件中存储了这个 module 对第三方依赖的全部信息

- 一个 Go Module 是一个 Go 包的集合. module 是有版本的
- 在 Go Module 模式下, 通常一个代码仓库对应一个 Go Module, 一个 Go Module 的顶层目录下会放置一个 go.mod 文件，每个 go.mod 文件会定义唯一一个 module
- go.mod 文件所在的顶层目录也被称为 module 的根目录

```shell
# 使用 go mod init 在当前目录创建一个 go.mod 文件
# go.mod 文件将当前项目变为一个 Go Module
go mod init blog  # 用前端类比, npm init -y, 创建一个 package.json

# go mod tidy 命令会扫描 Go 源码，并自动找出项目依赖的外部 Go Module 以及版本，
# 下载这些依赖并更新本地的 go.mod 文件
# 生成一个 go.sum 的文件记录直接依赖和间接依赖包的相关版本的 hash 值
go mod tidy # 前端 => npm install

# 特殊情况, 在 vendor 目录下，创建了一份这个项目的依赖包的副本
go mod vendor
```

```go
// go.mod
module blog

go 1.21.0

require github.com/sirupsen/logrus v1.8.1
```

上面我用了前端类比, 是为了更方便地理解, 当然了 go 和 node 包的管理方式完全不同

1. 为当前 module 添加一个依赖

```shell
# go get 命令将我们新增的依赖包下载到了本地 module 缓存里，
# 并在 go.mod 文件的 require 段中新增了一行内容
go get github.com/google/uuid
```

2. 升级/降级依赖的版本

Go Module 的版本号采用了语义版本规范，也就是版本号使用 vX.Y.Z 的格式。
其中 X 是主版本号，Y 为次版本号（minor），Z 为补丁版本号（patch）

以 `logrus` 为例

```shell
# 查询版本
go list -m -versions github.com/sirupsen/logrus

# 执行带有版本号的 go get 命令
# 选择 v1.7.0 版本的
go get github.com/sirupsen/logrus@v1.7.0

# 使用 go mod tidy
# 1. 先试用 go mode edit 明确告知依赖的版本
go mod edit -require=github.com/sirupsen/logrus@v1.7.0
go mod tidy


# 升级
go get github.com/sirupsen/logrus@v1.7.1
```

3. 添加一个主版本号大于 1 的依赖

如果新旧版本的包使用相同地导入路径，那么新包与旧包是兼容的

如果我们要为 Go 项目添加主版本号大于 1 的依赖，
我们就需要使用“语义导入版本”机制，在声明它的导入路径的基础上，加上版本号信息

```go
package main

import (
  _ "github.com/go-redis/redis/v7" // “_”为空导入
  "github.com/google/uuid"
  "github.com/sirupsen/logrus"
)

func main() {
  logrus.Println("hello, go module mode")
  logrus.Println(uuid.NewString())
}
```

```shell
# 加上版本号信息 v7
go get github.com/go-redis/redis/v7
```

4. 升级依赖版本到一个不兼容版本

语义导入版本的原则，不同主版本的包的导入路径是不同的

```go
import (
  _ "github.com/go-redis/redis/v8"
  "github.com/google/uuid"
  "github.com/sirupsen/logrus"
)
```

```shell
go get github.com/go-redis/redis/v8
```

5. 移除一个依赖

```shell
# 列出当前 module 的所有依赖
go list -m all

# 代码中删除 import 导入的依赖
# 自动分析源码依赖，而且将不再使用的依赖从 go.mod 和 go.sum 中移除
go mod tidy
```

### 入口函数和包初始化

1. main.main 函数: Go应用入口函数

```go
// 可执行程序的 main 包必须定义 main 函数，否则 Go 编译器会报错
package main

func main() {
    // 用户层执行逻辑
    ... ...
}
```

2. init 函数: Go包的初始化函数

进行包初始化的 `init` 函数

- 重置包级变量值
- 实现对包级变量的复杂初始化
- 在 init 函数实现 "注册模式"

```go
// init 函数的执行就都会发生在 main 函数之前
// init 函数也是一个无参数无返回值的函数
func init() {
    // 包初始化逻辑
    ... ...
}
```

Go 包可以拥有不止一个 init 函数，
每个组成 Go 包的 Go 源文件中，也可以定义多个 init 函数

先传递给 Go 编译器的源文件中的 init 函数，会先被执行；
而同一个源文件中的多个 init 函数，会按声明顺序依次执行

3. Go 包的初始化次序

- 依赖包按“深度优先”的次序进行初始化；
- 每个包内按以“常量 -> 变量 -> init 函数”的顺序进行初始化；
- 包内的多个 init 函数按出现次序进行自动调用

### 变量声明

| 内置原生类型                    | 默认值(零值) |
|---------------------------|---------|
| 所有整型类型                    | 0       |
| 浮点类型                      | 0.0     |
| 布尔类型                      | false   |
| 字符串类型                     | ""      |
| 指针 接口 切片 channel map和函数类型 | nil     |

- 包级变量: 包级别可见的变量。如果是导出变量（大写字母开头），那么这个包级变量也可以被视为全局变量
  - 只能使用带有 `var` 关键字的变量声明, 不能使用短变量声明形式
- 局部变量: Go 函数或方法体内声明的变量，仅在函数或方法体内可见


```go
// var 关键字
// a 变量名
// int 变量类型
// 10 变量的初始值
var a int = 10

var a int // a 的初始值为 int 类型的 零值: 0

// Go 提供了 变量声明块(block)语法形式
var (
	a int = 128
	b int8 = 6
	c string = "hello"
)

// 支持一行声明多个变量
var a, b, c = 5, 6,7
```

1. 省略类型信息的声明

```go
// Go 编译器会根据右侧变量初值自动推导出变量的类型
// 给这个变量赋予初值所对应的默认类型
var b = 13
// 显式类型转换
var b = int32(13)
```

2. 短变量声明

```go
// :=
a := 12
b := false

a, b := 12, false
```


3. 什么是变量遮蔽呢？

```go
var a = 11

func foo(n int) {
  // foo 函数中也使用了变量 a，但是 foo 函数中的变量 a 遮蔽了外面的包级变量 a
  a := 1
  a += n
}

func main() {
  fmt.Println("a =", a) // 11
  foo(5)
  fmt.Println("after calling foo, a =", a) // 11
}
```

变量遮蔽问题的根本原因，就是内层代码块中声明了一个与外层代码块同名且同类型的变量，
这样，内层代码块中的同名变量就会替代那个外层变量，参与此层代码块内的相关计算，
我们也就说内层变量遮蔽了外层同名变量


### 数值类型

**平台无关整型**

有符号整型: `int8 int16 int32 int64`
无符号整型: `uint8 uint16 uint32 uint64`

两者的区别在于最高二进制位(bit位)是否被解释为符号位
- 0 代表正 
- 1 代表负

**平台相关整型**

| 类型      | 32为长度      | 64位长度    |
|---------|------------|----------|
| int     | 32位(4字节)   | 64位(8字节) |
| uint    | 同上         | 同上       |
| uintptr | 存储任意一个指针的值 |          |

```go
var a, b = int(5), uint(6)
var p uintptr = 0x12345678
fmt.Println("signed integer a's length is", unsafe.Sizeof(a)) // 8
fmt.Println("unsigned integer b's length is", unsafe.Sizeof(b)) // 8
fmt.Println("uintptr's length is", unsafe.Sizeof(p)) // 8
```

**溢出问题**

无论上面哪种整型, 都有自己的取值范围, 超出范围, 发生整型溢出问题

**浮点型**

`float32 float64` 两种浮点类型, 默认值都是 0.0

**复数类型**

`complex64 complex128` 两种复数类型

```go
var c = 5 + 6i
var d = 0o123 + .12345E+5i // 83+12345i

var c = complex(5, 6) // 5 + 6i
var d = complex(0o123, .12345E+5) // 83+12345i

var c = complex(5, 6) // 5 + 6i
r := real(c) // 5.000000
i := imag(c) // 6.000000
```

**创建自己的数值类型**

通过 type 关键字基于原生数值类型来声明一个新类型

```go
type MyInt int32

var m int = 5
var n int32 = 6
var a MyInt = m // 错误：在赋值中不能将m（int类型）作为MyInt类型使用
var a MyInt = n // 错误：在赋值中不能将n（int32类型）作为MyInt类型使用
var a MyInt = MyInt(m) // ok
var a MyInt = MyInt(n) // ok
```

### 字符串类型

Go 语言源文件默认采用的是 Unicode 字符集，Unicode 字符集是目前市面上最流行的字符集，
它囊括了几乎所有主流非 ASCII 字符（包括中文字符）。
Go 字符串中的每个字符都是一个 Unicode 字符，并且这些 Unicode 字符是以 UTF-8 编码格式存储在内存当中的

字符串类型为 `string`

```go
"abc\n"
"中国人"
"\u4e2d\u56fd\u4eba" // 中国人
"\U00004e2d\U000056fd\U00004eba" // 中国人
"中\u56fd\u4eba" // 中国人，不同字符字面值形式混合在一起
"\xe4\xb8\xad\xe5\x9b\xbd\xe4\xba\xba" // 十六进制表示的字符串字面值：中国人
```

Go 语言中的字符串值也是一个`可空的字节序列`，
字节序列中的字节个数称为该字符串的长度。一个个的字节只是孤立数据，不表意

```go
var s = "中国人"
fmt.Printf("the length of s = %d\n", len(s)) // 9

for i := 0; i < len(s); i++ {
  fmt.Printf("0x%x ", s[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
}
fmt.Printf("\n")
```

字符串是由一个`可空的字符序列`构成

```go
var s = "中国人"
fmt.Println("the character count in s is", utf8.RuneCountInString(s)) // 3

for _, c := range s {
  // Unicode 字符
  fmt.Printf("0x%x ", c) // 0x4e2d 0x56fd 0x4eba
}
fmt.Printf("\n")
```

**rune 类型与字符字面值**

Go 使用 rune 这个类型来表示一个 Unicode 码点。
rune 本质上是 int32 类型的别名类型，它与 int32 类型是完全等价的

```go
type rune = int32
```
由于一个 Unicode 码点唯一对应一个 Unicode 字符, 所以
一个 rune 实例就是一个 Unicode 字符，一个 Go 字符串也可以被视为 rune 实例的集合。
我们可以通过字符字面值来初始化一个 rune 变量。

```go
// 通过单引号括起的字符字面值

'a'  // ASCII字符
'中' // Unicode字符集中的中文字符
'\n' // 换行字符
'\'' // 单引号字符
```

**Go字符串类型的内部表示**

string 类型其实是一个“描述符”，
它本身并不真正存储字符串数据，而仅是由一个指向底层存储的指针和字符串的长度字段组成的

```go
// $GOROOT/src/reflect/value.go

// StringHeader是一个string的运行时表示
type StringHeader struct {
    Data uintptr
    Len  int
}
```

**Go 字符串常见操作**

1. 下标操作
2. 字符迭代
3. 字符串连接
4. 字符串比较
   - Go 采用字典序的比较策略，分别从每个字符串的起始处，开始逐个字节地对两个字符串类型变量进行比较
5. 字符串转换
```go
var s = "中国人"
fmt.Printf("0x%x\n", s[0]) // 0xe4：字符“中” utf-8编码的第一个字节

// for 迭代
for i := 0; i < len(s); i++ {
  // index: 0, value: 0xe4 字符的 UTF-8 编码中的一个字节
  fmt.Printf("index: %d, value: 0x%x\n", i, s[i])
}
// for-range 迭代
for i, v := range s {
  // index: 0, value: 0x4e2d 字符串中 Unicode 字符的码点值
  fmt.Printf("index: %d, value: 0x%x\n", i, v)
}

s := "Rob Pike, "
s = s + "Robert Griesemer, "
s += " Ken Thompson"
fmt.Println(s) // Rob Pike, Robert Griesemer, Ken Thompson

func main() {
  // ==
  s1 := "世界和平"
  s2 := "世界" + "和平"
  fmt.Println(s1 == s2) // true

  // !=
  s1 = "Go"
  s2 = "C"
  fmt.Println(s1 != s2) // true

  // < and <=
  s1 = "12345"
  s2 = "23456"
  fmt.Println(s1 < s2)  // true
  fmt.Println(s1 <= s2) // true

  // > and >=
  s1 = "12345"
  s2 = "123"
  fmt.Println(s1 > s2)  // true
  fmt.Println(s1 >= s2) // true
}

var s string = "中国人"
// string -> []rune
rs := []rune(s)
fmt.Printf("%x\n", rs) // [4e2d 56fd 4eba]

// string -> []byte
bs := []byte(s)
fmt.Printf("%x\n", bs) // e4b8ade59bbde4baba

// []rune -> string
s1 := string(rs)
fmt.Println(s1) // 中国人

// []byte -> string
s2 := string(bs)
fmt.Println(s2) // 中国人
```

### 常量

```go
const Pi float64 = 3.14159265358979323846 // 单行常量声明

// 以const代码块形式声明常量
const (
    size int64 = 4096
    i, j, s = 13, 14, "bar" // 单行声明多个常量
)
```

**无类型常量**

即便两个类型拥有着相同的底层类型，
但它们仍然是不同的数据类型，不可以被相互比较或混在一个表达式中进行运算

```go
type myInt int
const n myInt = 13
const m int = n + 5 // 编译器报错：cannot use n + 5 (type myInt) as type int in const initializer

func main() {
    var a int = 5
    fmt.Println(a + n) // 编译器报错：invalid operation: a + n (mismatched types int and myInt)
}

// => 无类型常量
type myInt int
const n = 13

func main() {
  var a myInt = 5
  // 对于无类型常量参与的表达式求值，Go 编译器会根据上下文中的类型信息，
  // 把无类型常量自动转换为相应的类型后，再参与求值计算，这一转型动作是隐式进行的
  fmt.Println(a + n)  // 输出：18
}
```

**实现枚举**

```go
const (
    Apple, Banana = 11, 22 
    // 自动重复上一行
    Strawberry, Grape 
    Pear, Watermelon 
)

// iota
const (
  Apple, Banana = iota, iota + 10 // 0, 10 (iota = 0)
  Strawberry, Grape // 1, 11 (iota = 1)
  Pear, Watermelon  // 2, 12 (iota = 2)
)
```


### 定长数组到变长切片

1. 数组有那些基本特性

值类型

Go 语言的数组是一个长度固定的、由同构类型元素组成的连续序列

```go
// Go 编译器需要在编译阶段就知道数组类型的长度
// 长度为 N, 类型为 T
var arr [N]T

var arr1 [5]int
var arr2 [6]int
arr1 != arr2
```

数组类型不仅是逻辑上的连续序列, 在实际内存分配也占据一整块内存

```go
var arr = [6]int{1, 2, 3, 4, 5, 6}
fmt.Println("数组长度：", len(arr))           // 6
// 数组变量的总大小 64位平台 6 * 8
fmt.Println("数组大小：", unsafe.Sizeof(arr)) // 48

var arr1 [6]int // [0 0 0 0 0 0]

var arr2 = [6]int {
11, 12, 13, 14, 15, 16,
} // [11 12 13 14 15 16]

var arr3 = [...]int {
21, 22, 23,
} // [21 22 23]
fmt.Printf("%T\n", arr3) // [3]int
```

2. 多维数组

```go
var myArr [2][3][4]int
```

3. 切片

切片声明仅仅是少了一个“长度”属性

```go
var nums = []int{1, 2, 3, 4, 5, 6}
fmt.Println(len(nums)) // 6

// 动态地向切片中添加元素
nums = append(nums, 7) // 切片变为[1 2 3 4 5 6 7]
fmt.Println(len(nums)) // 7
```

4. Go是如何实现切片类型的

```go
// Go 编译器会自动为每个新创建的切片, 创建一个底层数组
type Slice struct {
	array unsafe.Pointer // 指向底层数组的指针
	len   int // 切片的长度, 当前元素的个数
	cap   int // 底层数组的长度, 切片的最大容量 cap > len
}

// 方法一: make 函数
s1 := make([]byte, 6, 10) // 10为cap值, 底层数组长度 6为去哦们的初始长度
s1 := make([]byte, 6) // cap = len = 6

// 方法二: array[low:high:max] 基于已存在的数组, 修改切片会修改数组
// max 默认省略, 默认值数组的长度
arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 起始元素从 low 所标识的下标值开始
// 切片的长度（len）是 high - low，它的容量是 max - low
sl := arr[3:7:9]
```

5. 切片的动态扩容

通过 append 操作向切片追加数据的时候，如果这时切片的 len 值和 cap 值是相等的，
也就是说切片底层数组已经没有空闲空间再来存储追加的值了，
Go 运行时就会对这个切片做扩容操作，来保证切片始终能存储下追加的新值

```go
var s []int
s = append(s, 11) 
fmt.Println(len(s), cap(s)) //1 1
s = append(s, 12) 
fmt.Println(len(s), cap(s)) //2 2
s = append(s, 13) 
fmt.Println(len(s), cap(s)) //3 4
s = append(s, 14) 
fmt.Println(len(s), cap(s)) //4 4
```


### 原生 map 类型

map 是 Go 语言提供的一种抽象数据类型，它表示`一组无序的键值对`

```go
// key 的类型必须支持“==”和“!=”两种比较操作符
// 函数类型、map 类型自身，以及切片类型是不能作为 map 的 key 类型的
map[key_type]value_type

map[string]int
map[int]int
```

map 变量的声明和初始化

```go
var m map[string]int // 默认值 nil

// map 类型，因为它内部实现的复杂性，无法“零值可用”
var m map[string]int // m = nil
// 发生运行时异常：panic: assignment to entry in nil map
m["key"] = 1

// 方法一: 使用字面量
m := map[int]string{} // Go 允许省略字面值中的元素类型
// 方法二: make函数为 map 类型变量进行初始化
m1 := make(map[int]string) // 未指定初始容量
m1 := make(map[int]string, 8) // 指定初始容量
```

map 基本操作

对同一 map 做多次遍历的时候，每次遍历元素的次序都不相同

map 引用类型

```go
// 插入新的键值对
m := make(map[int]string)
m[1] = "value1"
m[2] = "value2"
m[3] = "value3" // 相同key 新值会覆盖旧值

// 获取键值对数量
len(m)

// 查找和数据读取
m := make(map[string]int)
v, ok := m["key1"]
if !ok {
// "key1"不在map中
}
// "key1"在map中，v将被赋予"key1"键对应的value

// 删除
m := map[string]int {
"key1" : 1,
"key2" : 2,
}

fmt.Println(m) // map[key1:1 key2:2]
delete(m, "key2") // 删除"key2"
fmt.Println(m) // map[key1:1]
```


### 结构体

```go
// 定义一个空结构体
type Empty struct {} // 不包含任何字段的空结构体类型

var s Empty
println(unsafe.Sizeof(s)) // 0

// 使用其它结构体作为自定义结构体中字段的类型
type Person struct {
  Name string
  Phone string
  Addr string
}

type Book struct {
  Title string
  Author Person
  ... ...
}

var book Book
println(book.Author.Phone)

// 直接使用类型 嵌入字段
type Book struct {
  Title string
  Person
  ... ...
}
var book Book
println(book.Person.Phone) // 将类型名当作嵌入字段的名字
println(book.Phone)        // 支持直接访问嵌入字段所属类型中字段


// 结构体声明 初始化
type Book struct {}
var book Book
var book Book{}
book := Book{}

type Book struct {
  Title string              // 书名
  Pages int                 // 书的页数
  Indexes map[string]int    // 书的索引
}

// 使用复合字面值 按顺序依次给每个结构体字段进行赋值
var book = Book{"The Go", 700, make(map[string]int)}
var t = T{
  F2: "hello",
  F1: 11,
  F4: 14,
}
```

### 函数是 "一等公民"

```go
// 关键字 func
// 函数名 funcName
// param1...n 参数列表
// returnType 返回列表类型
func funcName(param1, param2, ...paramN) (returnType) {}

// 表述函数类型
func(io.Writer, string, ...interface{}) (int, error)

// 函数签名是相同的
func (a int, b string) (results []string, err error)
func (c int, d string) (sl []string, err error)
```

Go 语言中，函数参数传递采用是值传递的方式

**函数支持多返回值**

```go
func foo()                       // 无返回值
func foo() error                 // 仅有一个返回值
func foo() (int, string, error)  // 有2或2个以上返回值
```

**函数是一等公民**

- Go函数可以存储在变量中
- 支持在函数内创建并通过返回值返回
- 作为参数传入函数
- 拥有自己的类型

**结合多返回值进行错误处理**

error 类型与错误值构造

```go
// error 接口是 Go 原生内置的类型
type error interface {
	Error() string
}
```

难道为了构造一个错误值，我们还需要自定义一个新类型来实现 error 接口吗？

在标准库中提供了两种方便 Go 开发者构造错误值的方法： `errors.New` 和 `fmt.Errorf`

- 统一了错误类型
- 错误是值 可以像整型值那样对错误 `== !=` 逻辑比较
- 易扩展, 支持自定义错误上下文

```go
err := errors.New("demo error")
errWithCtx = fmt.Errorf("index %d", i)
```

**Go错误处理的策略**

- 透明错误处理策略
- "哨兵"错误处理策略
- 错误值类型检视策略
- 错误行为特征检视策略
  - 将某个包中的错误类型归类，统一提取出一些公共的错误行为特征，并将这些错误行为特征放入一个公开的接口类型中

```go
// 最常见的错误处理策略
err := doSomething()
if err != nil {
    // 不关心err变量底层错误值所携带的具体上下文信息
    // 执行简单错误处理逻辑并返回
    ... ...
    return err
}

func doSomething(...) error {
    ... ...
    return errors.New("some error occurred")
}

// 哨兵
data, err := b.Peek(1)
    if err != nil {
        switch err.Error() {
        case "bufio: negative count":
        // ... ...
        return
        case "bufio: buffer full":
        // ... ...
        return
        case "bufio: invalid use of UnreadByte":
        // ... ...
        return
        default:
        // ... ...
        return
    }
}

// $GOROOT/src/bufio/bufio.go
var (
    ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
    ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
    ErrBufferFull        = errors.New("bufio: buffer full")
    ErrNegativeCount     = errors.New("bufio: negative count")
)


```

**让函数更简洁健壮**

健壮三不要原则

- 不要相信任何外部输入的参数
- 不要忽略任何一个错误
- 不要假定异常不会发生

`panic`

Go 语言中，异常这个概念由 panic 表示

panic 指的是 Go 程序在运行时出现的一个异常情况。
如果异常出现了，但没有被捕获并恢复，Go 程序的执行就会被终止

无论在哪个 Goroutine 中发生未被恢复的 panic，整个程序都将崩溃退出

```go
panic("panic occurs in bar")
```

Go 也提供了捕捉 panic 并恢复程序正常执行秩序的方法，
我们可以通过 `recover 函数`来实现这一点

```go
func bar() {
    defer func() {
        if e := recover(); e != nil {
            fmt.Println("recover the panic:", e)
        }
    }()

    println("call bar")
    panic("panic occurs in bar")
    zoo()
    println("exit bar")
}
```

### 方法

- Go 中的方法必须是归属于一个类型的
- receiver 参数的基类型本身`不能为指针类型或接口类型`
- 方法声明要与 receiver 参数的基类型声明放在同一个包内
  - 我们不能为原生类型（诸如 int、float64、map 等）添加方法
  - 不能跨越 Go 包为其他包的类型声明新方法

```go
// 可以省略 func (T) M(t string)
// receiver 部分的参数名不能与方法参数列表中的形参名，以及具名返回值中的变量名存在冲突
func (t *T或T) MethodName(参数列表) (返回值列表) {
    // 方法体
}
```

- 如果 Go 方法要把对 receiver 参数代表的类型实例的修改，反映到原类型实例上，那么我们应该选择 `*T 作为 receiver` 参数的类型。
- 如果 T 类型需要实现某个接口，那我们就要使用 T 作为 receiver 参数的类型，来满足接口类型方法集合中的所有方法


**方法集合**

方法集合也是用来判断一个类型是否实现了某接口类型的唯一手段

方法集合决定接口实现的含义就是:

如果某类型 T 的方法集合与某接口类型的方法集合相同，
或者类型 T 的方法集合是接口类型 I 方法集合的超集，那么我们就说这个类型 T 实现了接口 I。
或者说，方法集合这个概念在 Go 语言中的主要用途，就是用来判断某个类型是否实现了某个接口。


**什么是类型嵌入**

支持两种类型嵌入:

- 接口类型的类型嵌入
- 结构体类型的类型嵌入

```go
type E interface {
  M1()
  M2()
}

type I interface {
  M1()
  M2()
  M3()
}

// 用接口类型 E 替代上面接口 I 定义中 M1 M2, 类型嵌入
type I interface {
  E
  M3()
}
```

```go
type T1 int
type t2 struct{
    n int
    m int
}

type I interface {
    M1()
}
// 结构体类型嵌入
type S1 struct {
    T1 // 标识符 T1 表示字段名为 T1，它的类型为自定义类型 T1
    *t2 // 标识符 t2 表示字段名为 t2，它的类型为自定义结构体类型 t2 的指针类型
    I // 标识符 I 表示字段名为 I，它的类型为接口类型 I
    a int
    b string
}
```

结构体类型的方法集合，包含嵌入的接口类型的方法集合

**结构体类型嵌入结构体类型**

```go
type T1 struct{}

func (T1) T1M1()   { println("T1's M1") } // T1 的方法集合包含：T1M1
func (*T1) PT1M2() { println("PT1's M2") } // *T1 的方法集合包含：T1M1、PT1M2；

type T2 struct{}

func (T2) T2M1()   { println("T2's M1") } // T2 的方法集合包含：T2M1
func (*T2) PT2M2() { println("PT2's M2") }// *T2 的方法集合包含：T2M1、PT2M2。

type T struct {
    T1
    *T2
}

func main() {
    t := T{
        T1: T1{},
        T2: &T2{},
    }

    dumpMethodSet(t) // 类型 T 的方法集合 = T1 的方法集合 + *T2 的方法集合
    dumpMethodSet(&t) // 类型 *T 的方法集合 = *T1 的方法集合 + *T2 的方法集合
}
```

**defined 类型与 alias 类型的方法集合**

凡通过类型声明语法声明的类型都被称为 defined 类型

基于类型别名（type alias）定义的新类型

```go
type I interface {
    M1()
    M2()
}
type T int
type NT T // 基于已存在的类型T创建新的defined类型NT
type NI I // 基于已存在的接口类型I创建新defined接口类型NI

type T struct{}

func (T) M1()  {}
func (*T) M2() {}

type T1 = T

func main() {
var t T
var pt *T
var t1 T1
var pt1 *T1

dumpMethodSet(t)
dumpMethodSet(t1)

dumpMethodSet(pt)
dumpMethodSet(pt1)
}
```