package models

type Article struct {
	Id            int
	Title         string
	ArticleCateId int // 外键
	State         int
	ArticleCate   ArticleCate `gorm:"foreignKey:ArticleCateId"` // 使用ArticleCateId为外键
}

func (Article) TableName() string {
	return "article"
}
