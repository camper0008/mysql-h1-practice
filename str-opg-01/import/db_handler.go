package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...
func DbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:example@/h1_opg01")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()

	// handle error
	if err != nil {
		panic(err)
	}

	// database object has  a method Close,
	// which is used to free the resource.
	// Free the resource when the function
	// is returned.

	return db
}
