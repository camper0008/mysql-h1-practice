package db_api

import (
	"csv-to-mysql/utils"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:example@/h1_opg01")
	utils.LogError(err)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	utils.LogError(err)

	return db
}
