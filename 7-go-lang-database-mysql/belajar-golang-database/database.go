package belajar_golang_database

import (
	"database/sql"
	"time"
)

// 146-database-pooling_test.go
// func GetConnection() *sql.DB {
// 	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")
// 	if err != nil {
// 		panic(err)
// 	}

// 	db.SetMaxIdleConns(10)
// 	db.SetMaxOpenConns(100)
// 	db.SetConnMaxIdleTime(5 * time.Minute)
// 	db.SetConnMaxLifetime(60 * time.Minute)

// 	return db
// }

// 149-tipe-data-column_test.go
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database?parseTime=true") // Terdapat params parseTime=true
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}