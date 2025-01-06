package main

import "fmt"

type Customer struct {
	// Biasanya orang membuat struct dengan kapital
	Name, Address string
	Age           int
}

func structWrapper() {
	// Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
	// Struct biasanya representasi data dalam program aplikasi yang kita buat
	// Data di struct disimpan dalam field
	// Sederhananya struct adalah kumpulan dari field

	// === Membuat data struct ===
	// struct adalah template data atau prototype data
	// struct tidak bisa langsung digunakan
	// namun kita bisa membuat data/object dari struct yang telah kita buat

	var person Customer
	person.Name = "Luqman"
	person.Address = "Indonesia"
	person.Age = 30

	fmt.Println(person)
	fmt.Println("name:", person.Name)

	// === Struct Literals ===
	// Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct
	person2 := Customer {
		Name: "Ola",
		Address: "Indonesia",
		Age: 28,
	}

	fmt.Println(person2)

	// Pada metode struct dibawah wajib sesuai urutan dan ukuran jumlah key dari struct harus sama
	person3 := Customer{"Joko", "Indonesia", 30}
	fmt.Println(person3)
}

// func main() {
// 	structWrapper()
// }
