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

---

### 接口

基本介绍

interface 类型可以定义一组方法, 但是这些不需要实现, 并且interface 不能包含任何变量.
到某个自定义类型(比如结构体)要使用的时候, 再根据具体情况把这些方法写出来

基本语法

```go
// 接口里的所有方法都没有方法体, 即接口的方法都是没有实现的方法, 体现了多态和高内聚低耦合的思想
// Golang中的接口, 不需要显式的视线, 只要一个变量, 含有接口类型的所有方法, 那么这个变量就实现了这个接口

type 接口名 interface {
	method1(参数列表) 返回值列表
	method2(参数列表) 返回值列表
	...
}

// 实现接口所有方法
func (t 自定义类型) method1(参数列表) 返回值列表  {
  // 方法实现
}
func (t 自定义类型) method2(参数列表) 返回值列表  {
  // 方法实现
}
// ...
```

案例

```go
// 接口 Usb 定义了两个方法
type Usb interface {
	// 声明两个未实现的方法
	Start()
	Stop()
}

// Phone 结构体
type Phone struct{}

type Camera struct{}

// 让 Phone 实现 usb 接口的方法, 实现了 Usb 接口
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

// 让 camera 实现 Usb 的方法, 实现了 usb 接口
func (c Camera) Start() {
	fmt.Println("相机开始工作")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

// 计算机
type Computer struct{}

// 编写一个方法 Working, 接收一个 Usb 接口类型变量
// 所谓实现了Usb接口, 就是指实现了 Usb 接口的所有方法
func (c Computer) Working(usb Usb) {
	// 通过 usb 接口变量来调用 Start Stop 方法
	usb.Start()
	usb.Stop()
}

func main() {
	c := Computer{}
	phone := Phone{}
	camera := Camera{}

	c.Working(phone)
	c.Working(camera)
}
```

**接口应用场景**

现在有一个项目经理, 管理三个程序员, 开发一个软件, 为了控制和管理软件, 项目经理
可以定义一些接口, 由程序员具体实现

```go
接口1 => 自定义类型1
接口1 => 自定义类型2

接口2 => 自定义类型3
接口2 => 自定义类型4
接口2 => 自定义类型5

接口4 => 自定义类型5
```
最佳实践


**接口注意事项和细节**

1. 接口本身不能创建实例, 但是可以指向一个实现了该接口的自定义类型的变量
2. 接口中的所有方法都没有方法体, 都是没有实现的方法
3. 在Golang中, 一个自定义类型需要将某个接口的所有方法都实现, 我们说这个自定义类型实现了该接口
4. 一个自定义类型只有实现了该接口, 才能将自定义类型的实例赋给接口类型
5. 只要是自定义数据类型, 就可以实现接口, 不仅仅是结构体类型
6. 一个自定义类型可以实现多个接口
7. Golang接口中不能有任何变量
8. 一个接口可以继承多个别的接口, 如果要实现A接口, 必须把B,C接口的方法页全部实现
9. interface类型默认是个指针(引用类型), 如果没有初始化interface就使用, 会输出 nil
10. 空接口 interface{}, 没有任何方法, 所以所有类型都实现了空接口, 可以把任何变量赋值给空接口

```go
// 5
type Aint interface {
	say()
}

type interger int

func (i interger) say()  {

}

type T interface{}
```

### 接口和继承

- 实现接口是对继承的补充
- 继承: 解决代码的`复用性和可维护性`
- 接口: `设计` 设计好各种规范, 让其它自定义类型去实现这些方法
- 接口比继承更加灵活
  - 继承是满足 is-a 的关系, 学生是人
  - 接口是满足 like-a 的关系, 猴子可以像鱼一样学会游泳
- 接口在一定程度上实现代码解耦


### 多态

变量具有多重形态, 在Go中, 多态特征是通过接口实现的. 可以按照统一的接口来调用不同的实现,
这时接口变量就呈现不同的形态

接口体现多态特征

- 多态参数
- 多态数组


### 类型断言

类型断言: 由于接口是一般类型, 不知道具体类型, 如果要转成具体类型, 需要使用类型断言
进行类型断言, 类型需匹配

语法 

```go
接口名称.(要转成类型的名称)
```

```go
type Point struct {
	x int
	y int
}

var a interface{}
var point Point = Point{1,2}
a = point // ok
// 如何将 a 赋给一个 Point变量?
var b Point
//b = a 不行
b = a.(Point)
fmt.Print(b)
```
























