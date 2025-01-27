package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)

/* === Auto Increment ===
- Kadang kita membuat sebuah table dengan id auto increment
- Dan kadang pula, kita ingin mengambil data id yang sudah kita insert ke dalam MySQL
- Sebenernya kita bisa melakukan query ulang ke database menggunakan SELECT LAST_INSERT_ID()
- Tapi untungnya di Golang ada cara yang lebih mudah
- Kita bisa menggunakan function (Result) LastInsertId() untuk mendapatkan Id terakhir yang
	dibuat secara auto increment
- Result adalah object yang dikembalikan ketika kita menggunakan function Exec
*/

func TestAutoIncrement(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "eko@gmail.com"
	comment := "Test koment"

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new comment with id ", insertId)
}