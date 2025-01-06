package main

import "fmt"

func comparationOperator() {
	// Operator Perbandingan
	// >, <, >=, <=, ==, !=

	var name1 = "Luqman"
	var name2 = "Budi"

	var result bool = name1 == name2
	fmt.Println(result)

	// Kalau komparasi string digunakan lebih dari maka akan melakukan komparasi byte
	

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
