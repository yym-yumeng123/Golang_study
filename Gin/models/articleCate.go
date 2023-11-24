package models

// 文章分类
type ArticleCate struct {
	Id    int // 主键
	Title string
	State int
}

func (ArticleCate) TableName() string {
	return "article_cate"
}
