package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func InitDB() {
	// 初始化数据库
	var err error
	db, err := gorm.Open("mysql", "mysql", "root:@tcp(127.0.0.1:3306)/summer_course?parseTime=true&charset=utf8&loc=Local")
	if err != nil {
		log.Panicf("Panic while connecting the gorm. Error: %s", err)
	}
	// 检查table是否存在，若不存在则新建table
	DB = db
	if !DB.HasTable(&Order{}) {
		DB.CreateTable(&Order{})
	}

	if !DB.HasTable(&Goods{}) {
		DB.CreateTable(&Goods{})
	}

}
