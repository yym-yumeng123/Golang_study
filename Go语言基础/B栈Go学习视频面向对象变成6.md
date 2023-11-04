### 面向对象编程应用实例

1. 声明/定义结构体, 确定结构体名
2. 编写结构体的字段
3. 编写结构体的方法

学生案例: Student结构体, 包含 name gender age id score, 有一个say方法, 返回 string类型, 方法返回信息中包含的所有字段值,

```go
type Student struct {
	Name string
	Gender string
	Age int
	id int
	score float64
}

func (s *Student) say() string {
  fmt.Println(s.Name, s.Age)
}

var s1 = Student{
	name: "tom",
	...
}
s1.say()
```

### 创建结构体变量时指定字段值

```go
// 1
var s Student = Student{"yym", 18}
s := Student{"yym", 10}
var s Student = Student{
	Name: "yym",
	Age: 18,
}

// 2
var s *Student = &Student{"yym", 18}
```

### 工厂模式

Golang的结构体没有构造函数, 使用工厂模式来解决

```go
type animal struct {
	Name string
	age  int
}

//因为animal结构体首字母小写,因此只能在 model 使用
// 通过工厂模式来解决

func NewAnimal(n string, s int) *animal {
	return &animal{
		Name: n,
		age:  s,
	}
}

// 字段是私有的
func (a *animal) GetAge() int {
	return a.age
}

```

### 面向对象编程思想-抽象

如何理解抽象

前面在定义一个结构体的时候, 实际上就是把一类事物的共有属性和行为提取出来, 形成一个物理模型. 这种研究问题的方法称为抽象

例如: 

银行账号1..n, 不管什么账号, 都有 账号, 密码, 余额..., 可以存款/取款/查询, 账号结构体 Account

```go
type Account struct {
	No string
	Pwd string
	Balance float64
}

// 存款
func (a Account) SaveMoney(money float64, pwd string) {
    if pwd == a.pwd {
			account.Balance += money
    }
}

// 取款
func (a Account) WithDraw(money float64, pwd string)  {

}

// 查询
```

### 面向对象三大特性

封装

1. 封装(encapsulation): 就是把抽象出的字段和对字段的操作封装在一起, 数据被保护在内部,程序的其它包只有通过被授权的操作, 才能对字段进行操作
2. 优点: 隐藏实现细节; 可以对数据进行验证, 保证安全合理
3. 体现封装: 对结构体中的属性进行封装; 通过方法包实现封装

封装的实现步骤

1. 将结构体, 字段的首字母小写, 其它包不能用
2. 给结构体所在包提供一个工厂模式的函数, 首字母大写, 类似一个构造函数
3. 提供一个首字母大写的 Set 方法, 用于对属性判断并赋值
4. 提供一个首字母大写的 Get 方法, 用于获取属性的值

```go
// 3.
func (val 结构体类型名) SetXxx(参数列表) 返回值列表{
	// 加入数据验证的业务逻辑
	val.字段 = 参数
}

// 4 
func (val 结构体类型名) GetXxx() {
	return val.字段
}
```

继承

- 继承可以解决代码复用
- 当多个结构体总在相同的属性和方法时, 可以从这些结构体中抽象出结构体, 在该结构体定义相同的属性和方法
- 其它的结构体不需要重新定义这些属性和方法, 只需嵌套一个 `匿名结构体`
- 也就是说: 在Golang中, 如果一个 struct 嵌套了另一个匿名结构体的字段和方法, 从而实现了继承特性

```go
// 匿名结构体语法
type Goods struct {
	Name string
}
type Book struct {
	Goods // 嵌套匿名结构体
}
var b Book
b.Goods.Name
b.Name

// 就近原则


// 结构体嵌多个匿名结构体, 两个匿名结构体有相同的字段和方法, 访问时, 需指定匿名结构体名字
type A struct {
  Name string
}
type B struct {
  Name string
}
type C struct {
	A
	B
}

var c C
c.A.Name
c.B.Name
```

一个 struct嵌套了一个有名结构体, 这种模式就是组合, 如果是组合关系, 那么在访问组合的结构体
的字段或方法时, 必须带上结构体的名字

```go
type A struct {
	Name string
	Age int
}

type C struct {
	a A
}

var c C
c.a.Name
```

嵌套匿名结构体后, 也可以在创建结构体变量时, 直接指定匿名结构体字段的值

```go
type Goods struct {
	Name string
	Price float64
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand
}
type TV1 struct {
	*Goods
	*Brand
}

tv := TV{Goods{"电视机", 5000.00}, Brand{"孩儿", "山东"}}
// &Good 实参的地址传给 指针类型
tv1 := TV1{&Good{"电视机", 5000.00}, &Brand{"孩儿", "山东"}}
```