package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//GetMysqlSession return db
func GetMysqlSession() *sql.DB {
	db, err := sql.Open("mysql", "connection")
	if err != nil {
		LogError(err)
	}
	err = db.Ping()
	if err != nil {
		LogError(err)
	}

	return db
}
