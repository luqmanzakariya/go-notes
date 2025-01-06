package main

import "fmt"

func switchExpression() {
	name := "Luqman"

	switch name {
	case "Luqman":
		fmt.Println("Halo Luqman")
	case "Ola":
		fmt.Println("Halo Ola")
	default:
		fmt.Println("Halo", name)
	}

	// Kode Program Switch dengan Short Statement
	switch length := len(name); length >= 5 {
	case true:
		fmt.Println("Jumlah karakter nama lebih atau sama dengan 5")
	case false:
		fmt.Println("Jumlah karakter kurang dari 5")
	}

	// Switch Tanpa Kondisi
	length2 := len(name)
	switch {
	case length2 > 10:
		fmt.Println("Jumlah karakter lebih dari 10")
	case length2 >= 5:
		fmt.Println("Jumlah karakter lebih dari 5")
	default:
		fmt.Println("Jumlah karakter kurang dari 5")
	}

}
