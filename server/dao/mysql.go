package dao

import (
	"database/sql"
	"time"
)

var DB *sql.DB

func InitMySql() {
	var err error
	dsn := "root:123456@tcp(192.168.31.229:3306)/home_data?charset=utf8mb4&parseTime=True"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	// 最大连接时长
	DB.SetConnMaxLifetime(time.Minute * 3)
	// 最大连接数
	DB.SetMaxOpenConns(10)
	// 空闲连接数
	DB.SetMaxIdleConns(10)
}
