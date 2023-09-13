#### if 语句

```go
if 布尔表达式 {
  /* 为 true 时执行 */
}

a := 10
if a < 20 {
  println("小于")
}



if 布尔表达式 {
  // true
} else {
  // false
}

if 布尔表达式1 {
  // 1 为 true 执行
  if 布尔表达式2 {
    // 2 为 true 执行
  }
}
```

#### switch 语句

```go
switch var1 {
  case val1:
    ...
  case val2:
    ...
  default:
    ...
}

marks := 90
switch marks {
  case 90: grade = "A"
  case 80: grade = "B"
  case 50,60,70 : grade = "C"
  default: grade = "D"  
}

switch {
  case grade == "A" :
      fmt.Printf("优秀!\n" )    
  case grade == "B", grade == "C" :
      fmt.Printf("良好\n" )      
  case grade == "D" :
      fmt.Printf("及格\n" )      
  case grade == "F":
      fmt.Printf("不及格\n" )
  default:
      fmt.Printf("差\n" );
}
```


#### for 循环

for 循环是一个循环控制结构，可以执行指定次数的循环

```go
// 语法

/**
 1. 先对表达式 1 赋初值
 2. 判别赋值表达式 init 是否满足给定条件，若其值为真，满足循环条件，则执行循环体内语句, 然后执行 post，进入第二次循环，再判别 condition；否则判断 condition 的值为假，不满足条件，就终止for循环，执行循环体外语句。
*/

/**
  init: 一般为赋值表达式, 给控制变量赋初值
  condition: 关系表达式或逻辑表达式， 循环控制条件
  post: 一般为赋值表达式，给控制变量增量或减量。
*/
for init; condition; post {  }

// for 循环的 range 可以对 slice map 数组 字符串等进行迭代循环
for key, value := range oldMap {
  newMap[key] = value
}

// 以上代码 key value可以省略, 如果只想读取key
for key := range oldMap


// 实例
for i := 0; i <= 10; i++ {
  sum += i
}

for sum <= 10{
  sum += sum
}

// 无限循环
sum := 0
for {
  sum++
}


strings := []string{"google", "runoob"}
for i, s := range strings {
  fmt.Println(i, s) // google runoob
}

numbers := [6]int{1, 2, 3, 5}
for i,x:= range numbers {
  fmt.Printf("第 %d 位 x 的值 = %d\n", i,x) // 1 2 3 5 0 0
}  


map1 := make(map[int]float32)
map1[1] = 1.0
map1[2] = 2.0
map1[3] = 3.0
map1[4] = 4.0

// 读取 key 和 value
for key, value := range map1 {
  fmt.Printf("key is: %d - value is: %f\n", key, value)
}

// 读取 key
for key := range map1 {
  fmt.Printf("key is: %d\n", key)
}

// 读取 value
for _, value := range map1 {
  fmt.Printf("value is: %f\n", value)
}
```