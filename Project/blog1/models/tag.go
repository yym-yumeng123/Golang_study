package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (*Tag) TableName() string {
	return "blog_tag"
}

func GetTags(pageNum, pageSize int, maps interface{}) (tags []Tag) {
	DB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// 在函数末端，我们已经显示声明了返回值，这个变量在函数体内也可以直接使用
func GetTotal(maps any) (count int64) {
	DB.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	DB.Select("id").Where("name=?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool {
	DB.Create(&Tag{
		Name:      name,
		CreatedBy: createdBy,
		State:     state,
	})
	return true
}

func EditTag(id int, data any) bool {
	DB.Model(&Tag{}).Where("id =?", id).Updates(data)
	return true
}

func DeleteTag(id int) bool {
	DB.Where("id=?", id).Delete(&Tag{})
	return true
}

func (u *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Set("CreatedOn", time.Now().Unix())
	return nil
}

func (u *Tag) BeforeUpdate(tx *gorm.DB) (err error) {
	tx.Set("ModifiedOn", time.Now().Unix())
	return nil
}
