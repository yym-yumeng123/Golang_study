### 反射

1. 反射可以在运行时动态获取变量的各种信息, 比如变量的类型(type), 类别(kind)
2. 如果时结构体变量, 还可以获取到结构体本身的信息(字段, 方法)
3. 通过反射, 可以修改变量的值, 可以调用关联的方法
4. 使用反射, import("reflect)

反射重要函数

1. `reflect.TypeOf(变量名)` 获取变量类型, 返回 `reflect.Type` 类型
2. `reflect.ValueOf(变量名)` 获取变量的值, 返回 `reflect.Value`类型, 是一个结构体类型
3. `变量, interface{} reflect.Value` 可以相互转换


```go
var stu Student
var num int

// 用于做反射
func test(b interface{}){
	// 如何将interface{} 转成 reflect.Value
	rVal := reflect.ValueOf(b)
	// 类型转回去
	iVal := rVal.Interface()
	// 如何将 interface{} 转成 原来的变量类型
	// 类型断言
	v := iVal.(Student)
}

test(stu)
```

注意事项和细节

1. `reflect.value.Kind` 获取变量的类别, 返回的是一个常量
2. Type 是类型, Kind 是类别, 可能相同, 可能也不同
3. 通过反射可以让变量在 interface{} 和 Reflect.Value 之间相互转换
4. 使用反射的方式来获取变量的值, 要求数据类型匹配, 比如 x 是 int, 应该使用 `reflect.Value().Int()`
5. 通过反射来修改变量, 注意当使用 `SetXXX`方法来设置需要通过对应指针类型来完成, 这样才能改变传入的值, 同时需要 `reflect.Value.Elem()` 方法
6. `reflect.Value.Elem()`

```go
func main() {
	// 通过反射修改值
	var num int = 10
	reflect1(&num)

	fmt.Println(num)
}

func reflect1(b interface{}) {
	rVal := reflect.ValueOf(b)
	rVal.Elem().SetInt(20)
}

```

反射最佳实践

1. 使用反射来遍历结构体的字段, 调用结构体的方法, 并获取结构体标签的值
   - `func (v Value) Method(i int) Value` 默认按方法名排序对应的 i 值, 从 0 开始
   - `func (v Value) Call(in []Value) []Value` 传入参数和返回参数是 []reflect.Value
