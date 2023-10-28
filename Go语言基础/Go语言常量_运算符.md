常量是一个简单值的标识符，在程序运行时，不会被修改的量

常量中的数据类型只可以是`布尔型、数字型（整数型、浮点型和复数）和字符串型`

```go
const identifier [type] = value

// 可以省略类型说明符 [type],
const b string = "abc"
const b = "abc" // 隐式类型定义

// 多个相同的类型
const c_name1, c_name2, c_name3 = value1, value2, value3
```

```go
// 常量
const LENGTH int = 10
const area int
const a, b, c = 1, false, "str" // 多重赋值

// 还可以用作枚举
const (
  Unknown = 0
  Female = 1
  Male = 2
)


// len() cap() unsize.Sizeof() 函数计算表达式的值, 常量表达式, 函数必须是内置函数

package main

import "unsafe"
const (
  a = "abc"
  b = len(a)
  c = unsafe.Sizeof(a)
)

func main(){
  println(a, b, c) // abc 3 16
}
```

#### iota

`iota`特殊常量, 认为是一个可以被编译器修改的常量

iota 在 const 关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

```go
// 用作枚举值
const (
  a = iota // 0
  b = iota // 1
  c = iota // 2
)

const (
  a = iota
  b
  c
)


// iota 用法
package main

import "fmt"

func main() {
  const (
      a = iota  //0
      b     //1
      c     //2
      d = "ha"  //独立值，iota += 1
      e     //"ha"  iota += 1
      f = 100  //iota +=1
      g     //100 iota +=1
      h = iota  //7,恢复计数
      i     //8
  )
  fmt.Println(a,b,c,d,e,f,g,h,i)
}

```

### Go 语言运算符

#### 算术运算符

- `+ - * / % ++ --`
- 对于除号 /, 整数之间做除法, 只保留整数部分二舍弃小数部分. `x := 19 / 5 => 3`
- Golang自增自减只能当一个独立语言使用, 不能 `b:=a++`
- Golang的`++ --` 只能写在变量后面, 不能写在前面

```go
var a int = 21
var b int = 10
var c int

c = a + b
fmt.Printf("第一行 - c 的值为 %d\n", c ) // 31
c = a - b
fmt.Printf("第二行 - c 的值为 %d\n", c ) // 11
c = a * b
fmt.Printf("第三行 - c 的值为 %d\n", c ) // 210
c = a / b // 保留整数位
fmt.Printf("第四行 - c 的值为 %d\n", c ) // 2
c = a % b // 取余数 a % b = a - a / b * b
fmt.Printf("第五行 - c 的值为 %d\n", c ) // 1
a++ // a = a + 1
fmt.Printf("第六行 - a 的值为 %d\n", a ) // 22
a=21   // 为了方便测试，a 这里重新赋值为 21
a-- // a = a - 1
fmt.Printf("第七行 - a 的值为 %d\n", a ) // 20
```

#### 关系运算符

- `== != > < >= <=`
- 结果都是 true/false

```go
var a int = 21
var b int = 10

if( a == b ) {
  fmt.Printf("第一行 - a 等于 b\n" )
} else {
  fmt.Printf("第一行 - a 不等于 b\n" ) // a 不等于 b
}
if ( a < b ) {
  fmt.Printf("第二行 - a 小于 b\n" )
} else {
  fmt.Printf("第二行 - a 不小于 b\n" ) // a 不小于 b
}

if ( a > b ) {
  fmt.Printf("第三行 - a 大于 b\n" ) // a 大于 b
} else {
  fmt.Printf("第三行 - a 不大于 b\n" )
}
/* Lets change value of a and b */
a = 5
b = 20
if ( a <= b ) {
  fmt.Printf("第四行 - a 小于等于 b\n" ) // a 小于等于 b
}
if ( b >= a ) {
  fmt.Printf("第五行 - b 大于等于 a\n" ) // b 大于等于 a
}
```

#### 逻辑运算符

- 用于链接多个条件, 最终结果也是一个 bool
- && 逻辑与也叫短路与, 如果第一个条件为 false, 第二个不会判断
- || 也叫短路或, 如果第一个条件为true, 第二个不会判断

A 值为 True，B 值为 False

- `&& -> (A && B) false`
- `|| -> (A || B) true` 
- `! -> !(A && B) true`

#### 赋值运算符

- 将某个运算后的值, 赋给指定的变量
- 运算顺序从右向左
- 赋值运算符左边, 只能是变量, 右边是变量 表达式 常量值

- `= += -= *= /= %= <<= >>=  &= ^= |=`

#### 其他运算符

- & 返回变量存储地址  &a -> 将给出变量的实际地址
- * 指针变量 *a 是一个指针变量

```go
var a int = 4
var b int32
var c float32
var ptr *int

ptr = &a // ptr 包含了 a 变量的地址
```

#### 运算符优先级

由上至下代表优先级由高到低

1. * / % << >> & &^ 
2. + - | ^ 单目运算符
3. == != < <= > >=
4. &&
5. ||
6. 赋值运算符 = += *= ...
7. 逗号 .