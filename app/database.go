package app

import (
	"database/sql"
	"login-app/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:acumalaka@tcp(localhost:3306)/login_app_test_db")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
