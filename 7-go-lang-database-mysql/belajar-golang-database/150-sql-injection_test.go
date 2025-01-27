package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)

/* === SQL Dengan Parameter ===
- Saat membuat aplikasi, kita tidak mungkin akan melakukan hardcode perintah SQL di kode Golang kita
- Biasanya kita akan menerima input data dari user, lalu membuat perintah SQL dari input user, dan
	mengirimnya menggunakan perintah SQL
*/

/* === SQL Injection ===
- SQL injection adalah sebuah teknik yang menyalahgunakan sebuah celah keamanan yang terjadi dalam
	lapisan basis data sebuah aplikasi
- Biasa, SQL injection dilakukan dengan mengirim input dari user dengan perintah yang salah, sehingga
	menyebabkan hasil SQL yang kita buat menjadi tidak valid
- SQL injection sangat berbahaya, jika sampai kita salah membuat SQL, bisa jadi data kita tidak aman
*/

/* === Solusinya? ===
- Jangan membuat query SQL secara manual dengan menggabungkan String secara bulat-bulat
- Jika kita membutuhkan parameter ketika membuat SQL, kita bisa menggunakan function Execute atau
	Query dengan parameter yang akan kita bahas di chapter selanjutnya
*/

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// username := "admin"
	// password := "admin"
	username := "admin'; #" // Contoh SQL Injection
	password := "admin"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	fmt.Println(script) // Cek Query String SQL Injection

	rows, err := db.QueryContext(ctx, script)
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
