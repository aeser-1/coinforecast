package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBconn() *sql.DB {

	db, err := sql.Open("mysql", "selectuser:Password.1234@tcp(localhost:3306)/project")
	if err != nil {
		panic(err.Error())
	}
	return db
}
