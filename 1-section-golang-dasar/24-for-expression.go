package main

import "fmt"

func forLoops() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke: ", counter)
		counter++
	}

	// For dengan statement, ada 2 statement
	// init statement: statement sebelum for dieksekusi
	// post statement: statement yang akan selalu dieksekusi di akhir tiap perulangan

	for i := 1; i <= 10; i++ {
		fmt.Println("i ke:", i)
	}

	slice := []string{"Eko", "Kurniawan", "Khannedy", "Rudy", "Joko"}

	for i := 0; i < len(slice); i++ {
		fmt.Println("Slice ke: ", i, slice[i])
	}

	// For Range
	// For bisa digunakan untuk melakukan iterasi terhadap semua data collection
	// Data collection contohnya Array, Slice, dan Map

	// Kode Program For Range
	for i, value := range slice {
		fmt.Println("index: ", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println("key: ", key, "=", value)
	}
}