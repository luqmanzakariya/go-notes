package main

import "fmt"

func mapWrapper() {
	// Pada array atau slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
	// Map aah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
	// Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
	// Berbeda dengan Array dan slice, jumlah data yang kita masukkan kedalam Map boleh sebanyak-banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])

	// Function Map
	// len(map) >> untuk mendapatkan jumlah data di map
	// map[key] >> Mengambil data di map dengan key
	// map[key] = value >> Mengubah data di map dengan key
	// make(map[TypeKey]TypeValue) = Membuat map baru
	// delete(map, key) = Menghapus data di map dengan key

	fmt.Println("=== Book ===")
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
