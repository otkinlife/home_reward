package base

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"home-reward/server/config"
)

var DB *sql.DB

func InitMySQL() {
	var err error
	DB, err = sql.Open("mysql", config.GlobalConfig.MySQLConnectInfo)
	if err != nil {
		panic(err)
	}
}
