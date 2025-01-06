package main

import "fmt"

func ifExpression() {
	name := "Joko"

	if name == "Luqman" {
		fmt.Println("Halo Luqman")
	} else if name == "Ola" {
		fmt.Println("Halo Ola")
	} else {
		fmt.Println("Halo", name, ",kamu bukan circle")
	}

	// === Kode Program If Short Statement ===

	// var length = len(name)
	// if length > 5 {
	// 	fmt.Println("Nama kamu lebih atau sama dengan 5")
	// } else {
	// 	fmt.Println("Nama kamu dibawah dari 5")
	// }

	// Menjadi
	if length := len(name); length > 5 {
		fmt.Println("Nama kamu lebih atau sama dengan 5")
	} else {
		fmt.Println("Nama kamu dibawah dari 5")
	}
}
