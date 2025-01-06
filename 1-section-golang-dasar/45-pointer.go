package main

// import "fmt"

// import "fmt"

/* === Pass by Value ===
- Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
- Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain,
  sebenarnya yang dikirim adalah duplikasi valuenya
*/

// type Address struct {
// 	City, Province, Country string
// }

// func main () {
// 	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
// 	address2 := address1 // copy value
	
// 	address2.City = "Bandung"

// 	fmt.Println(address1)
// 	fmt.Println(address2)
// }

/* === Pointer ===
- Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
- Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference
*/

/* === Operator & ===
- Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & diikuti dengan
  nama variable nya
*/

// func main () {
// 	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
// 	address2 := &address1
	
// 	address2.City = "Bandung"

// 	fmt.Println(address1) // ikut berubah
// 	fmt.Println(address2) // berubah menjadi bandung
// }