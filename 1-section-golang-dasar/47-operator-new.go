package main

// import "fmt"

/* === Operator new ===
- Sebelumnya untuk membuat pointer dengan menggunakan operator &
- Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
- Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
*/

// type Address struct {
// 	City, Province, Country string
// }

// func main () {
// 	// var alamat1 *Address = &Address{}
// 	var alamat1 *Address = new(Address) // sama dengan &Address{}
// 	var alamat2 *Address = alamat1

// 	alamat2.Country = "Indonesia"

// 	fmt.Println(alamat1)
// 	fmt.Println(alamat2)
// }