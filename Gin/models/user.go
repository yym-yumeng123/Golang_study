package models

// 结构体的名称首字母大写, 和数据库名称对应,  表名 user => 结构体 User, 表名 article_cate, 结构体 ArticleCate
// 默认情况表名是结构体名称的复数, 结构体 User, 表示这个模型默认操作的是 users 表
type User struct {
	Id       int // 字段名首字母大写 Id id
	Username string
	Age      int
	Email    string
	AddTime  int // add_time
}

func (u User) TableName() string {
	return "user"
}
