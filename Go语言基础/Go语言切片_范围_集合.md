### 切片 Slice

切片是对数组的抽象

数组的长度不可改变, 在特定场景中这样的集合就不太适用, Go 提供了("动态数组")与数组比, 切片的长度不是固定的, 可以追加元素, 在追加时使切片的容量增大

#### 定义切片

```go
// 声明一个未指定大小的数组来定义切片:
var identifier []type

// 切片不需要说明长度, 或使用 make() 函数来创建切片
var slice1 []type = make([]type, len)
slice1 := make([]type, len)

// 可以指定容量, 其中 capacity 为可选参数
make([]T, length, capacity)
```

切片初始化

```go
/**
[] 表示切片类型
{1,2,3} 初始化值一次是 1, 2, 3, 其 cap = len = 3
*/
s := [] int {1,2,3}

// 初始化切片 s, 是数组 arr 的引用
s := arr[:]

// 将 arr 中从下标startIndex到 endIndex - 1 下的元素创建为一个新的切片
s := arr[startIndex:endIndex]

// 默认 endIndex 时将表示一直到 arr 的最后一个元素
s := arr[startIndex:]

// 默认 startIndex 时将表示从 arr 的第一个元素开始
s := arr[:endIndex]

// 通过切片 s 初始化切片 s1
s1 := s[startIndex:endIndex]

// 内置函数 make() 初始化切片 s, []int 标识为其元素类型为 int 的切片
s := make([]int, len, cap)
```

#### len() cap() 函数

切片是可索引的, 可以由 len() 方法获取长度.
切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少

```go
package main
import "fmt"

func main() {
  var numbers = make([]int, 3, 5)
  printSlice(numbers)
}

func printSlice(x []int) {
  fmt.Println("len=%d cap=%d slice=%v\n", len(x), cap(x), x) // len=3 cap=5 slice=[0,0,0]
}
```

#### 空切片

一切切片在未初始化之前默认为 nil, 长度为 0

```go
package main
import "fmt"

func main() {
   var numbers []int
   printSlice(numbers)
   if(numbers == nil){
      fmt.Printf("切片是空的")
   }
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```

---

### 范围 Range

`range` 关键字用于 for 循环中`迭代数组(array)、切片(slice)、通道(channel)或集合(map)`的元素, 数组和切片中返回元素的索引和索引对应的值, 在集合中返回 `key-value` 对

```go
// 对 slice map 数组 字符串 进行迭代循环
for key, value := range oldMap {
  newMap[key] = value
}

// 只想读取 key
for key := range pldMap
for key, _ := range oldMap

// 只想读取 value
for _, value := range oldMap


// 实例: 2**%d 为 2对硬的次方
package main
import "fmt"

var pow = []int{1,2,4,8,16,32}

func main() {
  for i, v := range pow {
    fmt.Printft("2**%d = %d\n", i, v) // 2**0 =1 2**1 = 2 ...
  }
}

func main() {
  map1 := make(map[int]float32)
  map[1] = 1.0
  map[2] = 2.0
  map[3] = 3.0
  map[4] = 4.0

// 读取 key value
  for key, value := range map1 {
    fmt.Printf("key is: %d - value is: %f\n", key, value)
  }

  // 读取 key
  for key := range map1 {
    fmt.Printf("key is: %d\n", key)
  }

  for _, value := range map1 {
    fmt.Printf("value is: %f\n", value)
  } 
}


func main() {
  nums := []int{2,3,4}
  sum := 0
  for _, num := range nums {
    sum += num
  }

  kvs := map[string]{"a", "apple", "b", "banana"}
  for k, v := range kvs {

  } 

  for i, c := range "go" {

  }
}
```


### Map 集合

Map 是一种无序的键值对集合, 通过 key 来快速检索数据, key类似索引, 指向数据的值

Map 是一种集合, 可以迭代, 不过, map 是无序的, 

在获取 Map 值时, 如果键不存在, 返回该类型的零值, 例如 int 类型的零值是 0, string类型零值是 ""

#### 定义 Map

可以使用内建函数 make 或使用 map 关键字来定义 Map:

```go
/**
keyType 键的类型
ValueType 值的类型
initialCapacity 可选参数, 用于指定Map的初始容量, 
Map 的容量是指 Map中可以保存的键值对的数量, 当 Map 中的键值对数量达到容量时，Map 会自动扩容
*/
map_variable := make(map[KeyType]ValueType, initialCapacity)

// 创建一个控的 Map
m := make(map[string]int)

// 创建一个初始容量为 10 的 Map
m := make(map[string]int, 10)

// 字面量创建 Map: 
m := map[string]int {
  "apple": 1
  " banana": 2,
  "orange": 3
}

// 获取元素
v1 := m["apple"]
v2, ok := m["pear"] // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值

// 修改元素
m["apple"] = 5

// 获取 Map 的长度
len := len(m)

// 遍历 Map
for k, v := range m {
  fmt.Printf("key=%s, value=%d\n", k, v)
}

// 删除元素
delete(m, "banana")
```


#### delete() 函数

delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key