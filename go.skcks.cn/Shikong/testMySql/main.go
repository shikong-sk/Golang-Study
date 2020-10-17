package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type departmentModel struct {
	departmentId   string
	departmentName string
	active         int8
}

func main() {
	// 使用 go get github.com/go-sql-driver/mysql 安装 MySql 驱动
	db, err := sql.Open("mysql", "root:12341234@tcp(127.0.0.1:3306)/management_system?charset=utf8")
	fmt.Println(db, "\n", &db)
	if err != nil {
		fmt.Println(err)
	}
	// 最大连接数
	db.SetMaxOpenConns(100)
	// 闲置连接数
	db.SetConnMaxIdleTime(20)
	// 最大连接周期
	db.SetConnMaxLifetime(60 * time.Second)

	fmt.Println("检测数据库连接是否存活", db.Ping() == nil)

	var data departmentModel
	rows, err := db.Query("SELECT * FROM `ms_department`")
	if err != nil {
		fmt.Println(err)
	}
	// 获取字段名
	columns, _ := rows.Columns()
	// 获取字段类型
	columnsType, _ := rows.ColumnTypes()
	fmt.Println(columns)
	for _, v := range columnsType {
		fmt.Println(*v)
	}
	for rows.Next() {
		// 使用切片读取结果值
		d := make([]interface{}, len(columns))
		_ = rows.Scan(&d[0], &d[1], &d[2])
		fmt.Println(d)

		for _, v := range d {
			fmt.Printf("%v \t", string(v.([]uint8)))
		}
		fmt.Println()

		// 使用结构体读取结果值
		_ = rows.Scan(&data.departmentId, &data.departmentName, &data.active)
		fmt.Println(data)
	}
	_ = db.Close()
	fmt.Println("检测数据库连接是否存活", db.Ping() == nil)
}
