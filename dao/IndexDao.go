package dao

import (
	"fmt"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"task/setting"
)

/**
声明数据库全局连接对象
 */
var (
	DB *gorm.DB
)

/**
初始化 sqlite3 数据库连接
 */
func InitSqlite(cfg *setting.SQLite3Config) (err error) {
	// 获取数据库连接对象
	DB, err = gorm.Open(sqlite.Open(cfg.DB), &gorm.Config{})
	if err != nil {
		fmt.Printf("failed init the sqlite3...")
		return
	}
	return DB.Error
}

/**
释放数据库连接资源
 */
func Close() {
	// DB.Close()
}

/*
错误检查
 */
func checkErr(err error) {
	if err != nil { panic(err) }
}

/**
创建数据库
 */
func createDB() {
	// 创建数据表的 sql
	sqlTable := `
	CREATE TABLE IF NOT EXISTS todos(
		id INT(11) PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(255) NULL,
		status INT(1) NULL
	);
	`
	// 执行创建表的 sql 语句
	result := DB.Exec(sqlTable)
	fmt.Println("create database successfully and the affected row is:", result.RowsAffected)
}
