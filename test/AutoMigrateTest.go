package main

//import (
//	"fmt"
//	"gorm.io/driver/sqlite"
//	"gorm.io/gorm"
//)
//
///**
//定义表结构 Model
// */
//type Todos struct {
//	// gorm.Model
//	ID int `json:"id"`
//	Title string `json:"title"`
//	Status bool `json:"status"`
//}
//
///**
//测试 AutoMigrate 功能
// */
//func main() {
//	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
//	if err != nil {
//		return
//	}
//	err = db.AutoMigrate(&Todos{})
//	if err != nil {
//		return
//	}
//	fmt.Printf("AutoMigrate : ok")
//}
