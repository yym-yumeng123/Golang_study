### map

map 是 `key-value`数据结构, 又称为字段或者关联数组. 类似其他变成语言的集合

基本语法: `var map变量名 map[keytype]valueType`

- `keyType:` golang 中的map的key可以是很多类型: 比如 `bool 数字 string 指针 channel 接口 结构体 数组`, 通常为 `int string`
- slice map 还有 function 不可以, 因为没法 == 来判断
- `valueType: `和key基本一样, 通常为 `数字 string map struct`

```go
// 声明: 不会分配内存, 初始化需要 make, 分配内存后才能赋值和使用
var a map[string]string
var a map[string]int
var a map[int]string
var a map[string]map[string]string


// map 在使用前一定要make
// key 不能重复, 重复了会覆盖
// value 可以相同
var b map[string]string = make(map[string]string, 10)
b["name"] = "张三"
b["name1"] = "李四"
b["name1"] = "张三"
fmt.Println(b) // map[name:张三]

// make 使用
// 1. 先声明再赋值
var a map[string]int
a = make(map[string]int, 10)

// 2. 声明,就直接make
var a = make(map[string]int)

// 3. 声明直接赋值
var a = map[string]int{"name": 12}
```

**map的增删改查操作**

```go
// 增加和更新
map["key"] = value // key还没有, 就是增加, key存在就是修改

// 删除
delete(map, "key")
// 删除所有: 1. 遍历,逐个删除 2. 直接make一个新的空间
var a = make(map[string]int)

// 查找
val, findRes = map["key"]
```

- map遍历: `for-range`
- map的长度: len(map)
- map切片: 切片的数据类型如果是map, 择我们称为 slice of map, map切面, 这样使用map个数就可以动态变化了

**map使用细节**

- map是引用类型, 组训引用类型传递的机制, 在一个函数接受map,修改,会直接修改原来的map
- map的容量达到后,再想增加元素, 会自动扩容, 并不会panic
- map 的 value也经常使用 struct 类型, 更适合管理复杂的数据



