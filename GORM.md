### 什么是 ORM

`ORM(Object-Relational-Mapping)` 是在关系型数据库和对象之间做一个映射, 在操作数据库时
就不需要和复杂的 `sql` 语句打交道, 而是像操作对象一样操作sql就可以

`GORM` 是 `go` 语言的一个 `orm` 框架

```mysql
# 使用 mysql 创建一个 users 表
create table users
(
    id       int(11),
    username varchar(255),
    age      int(3),
    sex      int
);
```

```go
// 使用 gorm 对象创建一个 users 表
type User struct {
	ID int
	UserName string
	Age int
	Sex string
}

DB.CreateTable(&User{}) // 创建表
```

#### 安装 

```shell
# 安装 gorm
go get -u gorm.io/gorm 
# 安装 mysql 驱动
go get gorm.io/driver/mysql
```


#### 连接数据库 mysql

```go
var (
	DB  *gorm.DB
	err error
)

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 数据库用户名:密码@tcp(主机:端口)/gin?
	// charset=utf8mb4 设置字符集
	// parseTime=True 处理 time.Time 
	// loc=Local 时区设置,与本地时区保持一致
	dsn := "root:123456@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
    // Open 第一个参数: 数据库
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // 事务配置
	})
	if err != nil {
		fmt.Println(err)
	}
}
```

#### orm 操作表

```go
type User struct {
	ID int
	UserName string
	Age int
}

// 自定义表名
func (User) TableName() string{
	return "my_user"
}

// 创建表 表名为 结构体 小写复数 users
DB.CreateTable(&User{})
// 创建自定义名字的 表
DB.Table("user1").CteateTable(&User{})

// 删除表
DB.DropTable(&User{})
DB.DropTable("user1")

// 判断表是否存在
DB.HasTable(&USer{})

// 给表重命名
DB.RenameTable(oldName, newName)
```

#### 对表里面的数据 CRUD

```go
// 创建 users 表
DB.CreateTable(&User{})

// 新增 data
DB.Create(&User{UserName: "yym", Age: 18})

// 查询 data
DB.First(&user) // select * from users order by id limit 1;
DB.Take(&user) // select * from users limit 1;
DB.Last(&user) // select * from users order by id desc limit 1;
DB.First(&user, 10) // select * from users where id = 10;
DB.Find(&users, []int{1,2,3}) // select * from users where id in (1,2,3);
DB.First(&user, "id=?", "2121_aaa") // select * from users where id = "2121_aaa";
var user = User{ID: 10}
db.First(&user) // select * from users where id = 10
var result User
db.Model(User{ID: 10}).First(&result)  // SELECT * FROM users WHERE id = 10;

DB.Find(&users) // select * from users;

DB.Where("name=?", "yym").First(&user) // select * from users where name = 'yym' order by id limit 1 
DB.Where("name like ?", "%jin%").Find(&users) // select * from users where name like '%jin';
DB.Where("name=? and age>?", "yym", "10").Find(&user) // select * from users where name = 'yym' and age > 10;

// 更新 data, 先查询, 再更新
DB.First(&user)
user.Name = "yym2"
user.Age = 100
DB.Save(&user) // update users set name = 'yym2', age=100 where id=1;

// 更新单个字段
DB.Model(&user).Updata("age", 20) // update users set age=20 where id=1;
DB.Model(&user).Where("active=?", true).Update("name", "hello") // update users set name='hello' where id=1 and active=true;
// 更新多个字段, struct map[string]interface{}
DB.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false}) // update users set name='hello', age=18 where id =1;

// 删除 data
DB.Delete(&user) // delete from users where id = 1
DB.Where("name=?", "yym").Delete(&user)// delete from users where id =1 and name = 'yym';
DB.Delete(&Users{}, []int{1,2,3}) // delete from users where id in (1,2,3)
```

#### 结构体和表名的映射

1. 结构体没有驼峰命名, 表名就是: 结构体名小写+复数 => User 对应表 users
2. 有驼峰命名, 表名: 大写变小写前面加下划线,最后复数 => UserInfo 对应表 user_infos
3. 有连续大写字母, 表名: 连续大写字母变小写, 驼峰前加下划线小写复数 => DBUserInfo 对应 表 db_user_infos

```go
// 对应表名 users
type User struct {}

// 对应表名 user_infos
type UserInfo struct {}

// 对应表名 db_user_infos
type DBUserInfo struct {}
```

#### gorm.Model

GORM 定义一个 `gorm.Model` 结构体，其包括字段 `ID、CreatedAt、UpdatedAt、DeletedAt`

```go
// gorm.Model 的定义
type Model struct {
    ID        uint           `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
    gorm.Model
    Name string
}
// 等效于
type User struct {
    ID        uint           `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
    Name string
}
```

#### 字段标签

声明 `model` 时，`tag` 是可选的，GORM 支持以下 tag： tag 名大小写不敏感，但建议使用 camelCase 风格

- "primaryKey" 将列定义为主键
- "autoIncrement" 自增
- "-" 忽略该字段
- "not null" 指定列为 NOT NULL
- "index" 根据参数创建索引
- "unique" 唯一索引, 不能重复
- "column" 指定列名
- "size" 设置默认长度


```go
type Student struct {
	ID int `gorm:"primaryKey;autoIncrement"` // 主键 自动增长
	Name string `gorm:"index:idx_name;column:name1"` // 自定义索引名称
	Desc string `gorm:"-;size:10"` // 忽略不映射这个字段
}
```


#### 表关系

**一对一**

```go
// 一对一: 一个用户有一个扩展信息, 一个扩展信息对应一个用户, 外键可以加在任意表中
// belongs_to : 关系和外键的指定在同一方

// 用户表
type User struct {
	UserId string `gorm:"primaryKey;AutoIncrement"`
	Name string // 姓名
	Age int // 年龄
}

// 用户信息表
// `gorm:"foreignKey:MyUserId;AssociationForeignKey:UserId"` 指定外键
type UserInfo struct {
	InfoId int `gorm: primaryKey;AutoIncrement"`
	Pic string // 图片
	Address string
	User User `gorm:"foreignKey:MyUserId;AssociationForeignKey:UserId"` // 关联关系
	MyUserId int  // 指定外键
}


// has_one: 关系和外键的指不在同一方
type User struct {
	UserId string `gorm:"primaryKey;AutoIncrement"`
	Name string // 姓名 
	Age int // 年龄
	UserInfo UserInfo `gorm:"foreignKey:MyUserId"`
}

// 
type UserInfo struct {
	InfoId int `gorm: primaryKey;AutoIncrement"`
	Pic string // 图片 
	Address string 
	MyUserId int  // 指定外键
}
```

**一对多**

```go
// 作者
type Author struct {
	AID int `gorm:"primaryKey"`
	Name string
	Age int
	Article []Article // 关联关系
}
// 文章
type Article struct {
	ArID int `gorm:"primaryKey"`
	Title string
	Content string
	AID uint // 外键
	
}

// 一对多, 一个作者可以有多个文章, 一片文章属于一个作者
// 在文章表中加一个 所属作者 的 外键
```

**多对多**

```go
// 学生
type Student struct {
	SId int `gorm:"primaryKey"`
	Name string
	// 关联表
	Course []Course `gorm:"many2many:student_courses;"`
}

// 课程
type Course struct {
	CId int `gorm:"primaryKey"`
	teacher string
}

// 多对多, 一个学生有有个课程, 每个课程有很多学生上
// 新建一张关联表 student_schedules: id student_id schedule_id
```

#### 表关联操作

**一对一关联操作**

```go
// add 关联添加操作
info1 := UserInfo{
	Pic: "/xxx",
	Address: "上海",
	User: User{
		Name: "yym",
		Age: 18
	}
}

DB.create(&info1)

// find 关联查询操作 关联关系在 UserInfo 中, 所以从 info 入手
var userinfo UserInfo // userinfo 是源模型, 主键不能为空
// 1. Association
DB.First(&userinfo, "info_id=?", 1) // 查不到关联关系
// Model参数: 要查询的表数据
// Association参数, 关联到具体的模型名 User
// Find参数: 查询的数据要放在什么字段中
DB.Model(&userinfo).Association("User").Find(&userinfo.User)
// 带的条件的查询
codes := []string{"zh-CN", "en-US", "ja-JP"}
DB.Model(&user).Where("code IN ?", codes).Association("Languages").Find(&languages)

// 2. Preload 方式 预加载
DB.Preload("User").Find(&userinfo, "info_id=?", 1)

// 3. Related 方式
DB.First(&userinfo, "info_id=?", 1)
var user User
// 通过 userinfo 查出来的 User 字段信息放入新的容器 User 中
DB.Model(&userinfo).Related(&user, "User")

// update 关联更新 通过 info 对 user 表数据更新
DB.Preload("User").Find(&userinfo "info_id", 1) // 查询
DB.Model(&userinfo.user).Update("age", 20) // 更新

// delete 关联删除
DB.Preload("User").Find(&userinfo "info_id", 1) // 查询
DB.Delete(&userinfo.User) // 删除userinfo 的 User 表中的记录
```

**一对多关联操作**

```go
// author article 两个表, 一个 author -> 多个 article, 操作 author 表

// add 关联添加
author := Author{
	Name: "yym",
	Age: 18,
	Article: []Atricle{
		{
			Title: "HTML",
			Content: "122"
		},
		{
			Content: "CSS"
			Content: "122"
		}
	},
}

DB.Create(&author)

// find 查询
var author Author
// 条件是 文章的条件
DB.Model(&author).Where("ar_id=?", 1).Association("Article").Find(&author.Article)

// update 更新
// 先查询, 再更新
DB.Preload("Article").Find(&author, "a_id=?", 1)
DB.Model(&author.Article).Where("ar_id=?", 1).Update("title", "JS入门")

// 删除
DB.Where("ar_id=?", 2).Delete(&author.Article)
```

**多对多关联操作**

```go
// student course两个表 关联关系在 student 中, 操作 student

// add
stu := Student{
	Name: "yym",
	Course: []Course{
		{
			teacher: "张三",
        },{
			teacher: "李四"
				}
	}
}

DB.Create(&stu)
```

#### 常用方法

```go
DB.First() // 按条件查询, 升序排列 查询出一条记录

// 有对应数据, 就查出来, 没有, 就创建
DB.Where(User{Name: "non_existing"}).Attrs(User{Email: "fake@fake.org"}).FirstOrCreate(&user)

DB.Last() // 按条件查询, 降序排列 查询出一条记录

DB.Take() // 按条件查询, 和 First 类似, 未排序

DB.Find(&user, 1) // select * from where id = 1

DB.Where() // 查询 加入指定条件 

DB.Select("name, age") // 对字段过滤

DB.Create() // 添加单条数据

DB.Save() // 保存更新数据

DB.Update() // 更新一列
DB.Updates() // 更新多列

DB.Delete() // 删除数据

DB.Not("user_id = ?", 1).Find(&users) // 排除 id = 1 的值
DB.Or("user_id=?", 2) // 多个条件的查询
DB.Order("age desc") // 升序或降序排列

DB.Limit(10) // 获取记录的最大数量
DB.Offset(1).Limit(3) // 跳过几条,偏移, 和 Limit 结合使用
// 结果扫描到另一个结构体, 可以对 JSON 进行处理
DB.Scan(&user2) // 把获取的数据扫描到 &user2 中, 结构体字段一致

DB.Count(&count) // 计数
// select age count(*) from users group by age
DB.Group("age") // 根据年龄进行分组
// select age, count(*) from users group by age having (age > 18)
DB.Having("age >18") // 分组以后进行过滤


// 左连接 右连接
// 左连接 users 表是全的, user_infos 表数据不全
//select * from users left join user_infos on users.id = user_infos.info_id

// 右连接 users 表不全, user_infos 表是全的
//select * from users right join user_infos on users.id = user_infos.info_id
type NewUserInfo struct {} // 新的结构体, 扫描数据进来
DB.Joins("left join user_infos on users.id = user_infos.info_id").Find(&users).Scan(&NewUserInfo)

DB.DeBug() // 打印当前行的 sql 语句

// 日志级别
logger.Default.LogMode(logger.Silent)

// 操作原生 sql 查询 Raw
DB.Raw("select * from users").Find(&users)
DB.Raw("select * from users where age = ?", 14).Find(&users)
// 操作原生sql Exec 增加 删除 修改
DB.Exec("insert into users (age, name) values (?, ?)", 33, "张五")
DB.Exec("delete from users where user_id = ?", 1)
DB.Exec("update users set name = ? where user_id = ?", "张三", 3)
```

#### 日志

Golang标准库的日志框架比较简单, 使用第三方日志 `logrus`

```shell
go get -u github.com/sirupsen/logrus
```