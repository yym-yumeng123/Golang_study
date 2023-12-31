package models

import (
	"gorm.io/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	Createdby  string `json:"created_by"`
	Modifiedby string `json:"modified_by"`
	State      int    `json:"state"`
}

func (a *Article) TableName() string {
	return "blog_article"
}

func (a *Article) BeforeCreate(scope *gorm.DB) error {
	scope.Set("CreatedOn", time.Now().Unix())

	return nil
}

func (a *Article) BeforeUpdate(scope *gorm.DB) error {
	scope.Set("ModifiedOn", time.Now().Unix())

	return nil
}

func ExistArticleByID(id int) bool {
	var article Article
	DB.Select("id").Where("id = ?", id).First(&article)

	if article.ID > 0 {
		return true
	}

	return false
}

func GetArticleTotal(maps interface{}) (count int64) {
	DB.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	DB.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	DB.Where("id = ?", id).First(&article)
	//DB.Model(&article).Related(&article.Tag)

	return
}

func EditArticle(id int, data interface{}) bool {
	DB.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	DB.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		Createdby: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	DB.Where("id = ?", id).Delete(Article{})

	return true
}
