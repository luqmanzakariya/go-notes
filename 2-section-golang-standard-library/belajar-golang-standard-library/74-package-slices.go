package main

// import (
// 	"fmt"
// 	"slices"
// )

/* === Package slices ===
- Di Golang versi terbaru, terdapat fitur bernama Generic, fitur ini akan kita bahas khusus dikelas Golang
	Generic
- Fitur Generic ini membuat kita bisa membuat parameter dengan tipe data yang bisa berubah-ubah, tanpa harus
	menggunakan interface kosong/any
- Salah satu package yang menggunakan fitur Generic ini adalah package slices
- Package slices ini digunakan untuk memanipulasi data di slice
*/

// func main() {
// 	names := []string{"John", "Paul", "George", "Ringo"}
// 	values := []int{100, 95, 80, 90}

// 	fmt.Println(slices.Min(values))
// 	fmt.Println(slices.Max(values))
// 	fmt.Println(slices.Contains(names, "Paul"))
// 	fmt.Println(slices.Index(names, "George"))
// }