package main

// import "fmt"

/* === Operator * ===
- Saat kita mengubah variable pointer, maka yang hanya berubah hanya variable tersebut.
- Semua variable yang mengacu ke data yang sama tidak akan berubah
- Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa
  menggunakan operator *
*/

// type Address struct {
// 	City, Province, Country string
// }

// func main () {
// 	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
// 	address2 := &address1
// 	address2.City = "Bandung"

// 	fmt.Println(address1) // ikut berubah
// 	fmt.Println(address2) // berubah menjadi bandung

// 	// Tanpa operator *
// 	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
// 	// Dengan operator *
// 	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
// 	fmt.Println(address1)
// 	fmt.Println(address2)
	
// }
