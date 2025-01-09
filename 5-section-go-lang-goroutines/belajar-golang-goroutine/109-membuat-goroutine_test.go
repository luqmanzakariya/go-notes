package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

/* === Membuat Goroutine ===
- Untuk membuat goroutine di Golang sangatlah sederhana
- Kita hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan
	dalam goroutine
- Saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara
	asynchronous, artinya tidak akan ditunggu sampai function tersebut selesai
- Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita
	buat selesai
*/

func RunHelloWorld(){
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T){
	go RunHelloWorld() // Tinggal ditambahkan go untuk menggunakan go routine
	fmt.Println("Ups")

	time.Sleep(1 * time.Second) // Apabila ini di comment maka ada kemungkinan hello world tidak berjalan karena program telah selesai
}