package main

//
//import (
//	"database/sql"
//	"fmt"
//	"time"
//
//	// if you got the error info: exec: "gcc": executable file not found in %PATH%
//	// you need to download the GCC and install it: https://sourceforge.net/projects/mingw-w64/
//	_ "github.com/mattn/go-sqlite3"
//)
//
///**
//连接 SQLite 数据库测试程序
// */
//func main() {
//
//	// 连接数据库,如果不存在则创建该数据库
//	db, err := sql.Open("sqlite3", "./todo.db")
//	checkErr(err)
//
//	// 当程序结束时关闭数据库连接
//	defer db.Close()
//
//	// 创建表sql
//	sql_table := `
//    CREATE TABLE IF NOT EXISTS userinfo(
//       uid INTEGER PRIMARY KEY AUTOINCREMENT,
//       username VARCHAR(64) NULL,
//       departname VARCHAR(64) NULL,
//       created DATE NULL
//    );
//    `
//	// 执行创建表的 sql 语句
//	db.Exec(sql_table)
//
//	// 插入数据
//	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?, ?, ?)")
//	checkErr(err)
//	res, err := stmt.Exec("GoogTech", "Google", "2020-11-13")
//	checkErr(err)
//	id, err := res.LastInsertId()
//	checkErr(err)
//	fmt.Println("inserted successfully and the last insert id: ", id)
//
//	// 更新数据
//	stmt, err = db.Prepare("update userinfo set username = ? where uid = ?")
//	checkErr(err)
//	res, err = stmt.Exec("GolandTech", 1)
//	checkErr(err)
//	affect, err := res.RowsAffected()
//	checkErr(err)
//	fmt.Println("updated successfully and the affect row is: ", affect)
//
//	// 查询数据
//	rows, err := db.Query("SELECT * FROM userinfo")
//	checkErr(err)
//	var uid int
//	var username string
//	var department string
//	var created time.Time
//	for rows.Next() {
//		err = rows.Scan(&uid, &username, &department, &created)
//		checkErr(err)
//		fmt.Println(uid)
//		fmt.Println(username)
//		fmt.Println(department)
//		fmt.Println(created)
//	}
//	rows.Close()
//
//	// 删除数据
//	stmt, err = db.Prepare("delete from userinfo where uid = ?")
//	checkErr(err)
//	res, err = stmt.Exec(1)
//	checkErr(err)
//	affect, err = res.RowsAffected()
//	checkErr(err)
//	fmt.Println("deleted successfully and the affect row is: ", affect)
//}
//
//// 错误检查
//func checkErr(err error) {
//	if err != nil { panic(err) }
//}
