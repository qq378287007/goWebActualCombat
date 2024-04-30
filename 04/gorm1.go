package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User2 用户信息
type User2 struct {
	//gorm.Model
	Id    int
	Name  string
	Phone string
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/chapter4?"+"charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// 自动迁移
	//db.AutoMigrate(&UserInfo{})

	// 创建记录
	u1 := &User2{4, "Jack", "18888888888"}
	db.Create(u1)
}
