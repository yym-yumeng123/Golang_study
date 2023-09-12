### Go 语言变量

变量可以通过变量名访问

Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字

```go
// 使用 var 关键字
var identifier type

// 一次声明多个变量
var identifier1, identifier2 type

var a string = "yym"
var b, c int = 1, 2
```

#### 变量声明

1. 指定变量类型, 如果没有初始化, 则变量默认位零值
   - 数值类型 `0`
   - 布尔 `false`
   - 字符串 `""`
   - 以下集中类型是 `nil`

```go
var v_name v_type
v_name = value

// 零值就是变量没有做初始化时系统默认设置的值。
var a = "RUNOOB"
var b int // 0
var c bool // flase


// nil
var a *int
var a []int
var a map[string] int
var a func(string) int
var a error // error 是接口
```

2. 根据值自行判定变量类型

```go
var  v_name = value

var d = true // true
```

3. 如果变量已经使用 var 声明过了，再使用 `:=` 声明变量，就产生编译错误，格式

```go
v_name := value


var intVal int
intVal := 1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明

intVal := 1 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
// => 相等于
var intVal int
intVa = 1


var f string = "yym"
// =>
f := "yym"
```

#### 多变量声明

```go
//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误


// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)


var x, y int
var (
  a int
  b bool
)
var c, d int = 1, 2
var e, f = 123, "hello"

func main(){
    g, h := 123, "hello"
    fmt.Println(x, y, a, b, c, d, e, f, g, h)
}
```

### 值类型和引用类型

所有像 `int、float、bool 和 string` 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值：

```go
int i -> 7 // 32 bit word
```

当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在`内存中将 i 的值进行了拷贝`

```go
int i -> 7
int j -> 7 // 拷贝
```

可以通过 `&i` 来获取变量 i 的内存地址, 例如：0xf840000040（每次的地址都可能不一样）。

内存地址会根据机器的不同而有所不同，甚至相同的程序在不同的机器上执行后也会有不同的内存地址。因为每台机器可能有不同的存储器布局，并且位置分配也可能不同

一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个字所在的位置

```go
ref r1 -> address1 -> value of r1
ref r2 -> address2 -> value of r1
```

这个内存称为指针, 这个指针实际上也被存在另外的某一个值中

同一个引用类型的指针指向的多个字可以是在连续的内存地址中（内存布局是连续的），这也是计算效率最高的一种存储形式；也可以将这些字分散存放在内存中，每个字都指示了下一个字所在的内存地址

当使用赋值语句 r2 = r1 时，只有引用（地址）被复制

如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容，在这个例子中，r2 也会受到影响


### 简短形式, 使用 := 赋值操作符

`只能被用在函数体内，而不可以用于全局变量的声明与赋值`

可以在变量的初始化时省略变量的类型而由系统自动推断，声明语句写上 var 关键字其实是显得有些多余了, 将它们简写为 `a := 50` 或 `b := false` (首选方式)

