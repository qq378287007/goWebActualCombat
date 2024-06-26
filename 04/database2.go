package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	//连接数据库
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed, err:%v\n", err)
		return
	}
}
