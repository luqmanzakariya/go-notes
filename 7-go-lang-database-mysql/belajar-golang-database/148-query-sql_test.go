package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)

/* === Query SQL ===
- Untuk operasi SQL yang tidak membutuhkan hasil, kita bisa menggunakan perintah Exec, namun jika
	kita membutuhkan result, seperti SELECT SQL, kita bisa menggunakan function yang berbeda
- Function untuk melakukan query ke database, bisa menggunakan function
	(DB) QueryContext(context, sql, params)
*/

/* === Rows ===
- Hasil Query function adalah sebuah data structs sql.Rows
- Rows digunakan untuk melakukan iterasi terhadap hasil dari query
- Kita bisa menggunakan function(Rows) Next() (boolean) untuk melakukan iterasi terhadap data hasil
	query, jika return data false, artinya sudah tidak ada data lagi didalam result
- Untuk membaca tiap data, kita bisa menggunakan (Rows) Scan(columns...)
- Dan jangan lupa, setelah menggunakan Rows, jangan lupa untuk menutupnya menggunakan
	(Rows) Close()
*/

func TestQuerySql(t *testing.T){
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
	}
}