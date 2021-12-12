package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/blog?parseTime=true")
	if err != nil {
		panic(err)
	}
	return db
}
