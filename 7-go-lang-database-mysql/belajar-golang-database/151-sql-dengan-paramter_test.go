package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)

/* === SQL Dengan Parameter ===
- Sekarang kita sudah tahu bahaya nya SQL injection jika menggabungkan string ketika membuat query
- Jika ada kebutuhan seperti itu, sebenarnya function Exec dan Query memiliki parameter tambahan yang
	bisa kita gunakan untuk mensubtitusi parameter dari function tersebut ke SQL query yang kita buat.
- Untuk menandai sebuah SQL membutuhkan parameter, kita bisa gunakan karakter ? (tanda tanya)
*/

/* === Contoh SQL ===
- SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1
- INSERT INTO user(username, password) VALUES(?, ?)
- Dan lain-lain
*/

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username := "admin"
	// password := "admin"
	username := "admin'; #" // Contoh SQL Injection
	password := "admin"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	fmt.Println(script) // Cek Query String SQL Injection

	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestExecSqlParameter(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username := "eko"
	// password := "eko"
	username := "eko', DROP TABLE user; #" // Contoh SQL Injection
	password := "eko"

	script := "INSERT INTO user(username, password) VALUES(?, ?)"

	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new user")
}