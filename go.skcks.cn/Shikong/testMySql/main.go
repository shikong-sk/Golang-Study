package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 使用 go get github.com/go-sql-driver/mysql 安装 MySql 驱动
	db, err := sql.Open("mysql", "root:12341234@tcp(127.0.0.1:3306)/management_system?charset=utf8")
	fmt.Println(db, "\n", &db, "\n", err)
}
