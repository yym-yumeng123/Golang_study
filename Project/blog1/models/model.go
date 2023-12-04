package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

type Model struct {
	ID         int `gorm:"primaryKey" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func Init() {
	var (
		err error
	)

	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	db, err := gorm.Open(mysql.Open("root:YYm1994@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	DB = db

	DB.AutoMigrate(&Article{}, &Tag{})

	if err != nil {
		log.Println(err)
	}

	//db.DB().SetMaxIdleConns(10)
	//db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
}
