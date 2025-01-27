package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

/* === Tipe Data Column ===
- Sebelumnya kita hanya membuat table dengan tipe data di kolom nya berupa VARCHAR
- Untuk VARCHAR di database, biasanya kita gunakan String di Golang
- Bagaimana dengan tipe data yang lain?
- Apa representasinya di Golang, misal tipe data timestamp, date dan lain-lain
*/

/* === Mapping Tipe Data ===
Tipe Data Database									Tipe Data Golang
VARCHAR, CHAR												string
INT, BIGNIT													int32, int64
FLOAT, DOUBLE												float32, float64
BOOLEAN															bool
DATE, DATETIME, TIME, TIMESTAMP			time.Time
*/

/* === Error Tipe Data Date ===
- Secara default, Driver MySQL untuk Golang akan melakukan query tipe data DATE, DATETIME, TIMESTAMP
	menjadi []byte / []uint8. Dimana ini bisa dikonversi menjadi String, lalu di parsing menjadi time.Time
- Namun hal ini merepotkan jika dilakukan manual, kita bisa meminta Driver MySQL untuk Golang secara
	otomatis melakukan parsing dengan menambah parameter parseDate=true
- Contoh: Cek file database.go
*/

/* === Nullable Type ===
- Golang database tidak mengerti dengan tipe data NULL di database
- Oleh karena itu, khusus untuk kolom yang bisa NULL di database, akan jadi masalah jika kita melakukan
	Scan secara bulat-bulat menggunakan tipe data representasinya di Golang
*/

/* === Error Data Null ===
- Konversi secara otomatis NULL tidak didukung oleh Driver MySQL Golang
- Oleh karena itu, khusus tipe kolom yang bisa NULL, kita perlu menggunakan tipe data yang ada dalam
	package sql
*/

/* === Error Data Null ===
Tipe Data Golang		Tipe Data Nullable
string							database/sql.NullString
bool								database/sql.NullBool
float64							database/sql.NullFloat64
int32								database/sql.NullInt32
int64								database/sql.NullInt64
time.Time						database/sql.NullTime
*/

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("===============================")
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("email: ", email.String)
		}
		fmt.Println("balance: ", balance)
		fmt.Println("rating: ", rating)
		if birthDate.Valid {
			fmt.Println("birtDate: ", birthDate.Time)
		}
		fmt.Println("married: ", married)
		fmt.Println("createdAt: ", createdAt)
	}
}
