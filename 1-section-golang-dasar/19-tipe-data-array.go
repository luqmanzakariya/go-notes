package main

import "fmt"

func array() {
	// index di array sama dengan javascript

	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Luqman"
	names[2] = "Zakariya"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Membuat Array Langsung
	var values = [3]int{
		90,98,100,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// Function array
	// len(array) (untuk mendapatkan panjang array)
	// array[index] (mendapat data di posisi index)
	// array[index] = value (mengubah data di posisi index)
	
	values[0] = 3
	fmt.Println(len(values))
	fmt.Println(values)
}