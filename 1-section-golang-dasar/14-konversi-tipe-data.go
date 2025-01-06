package main

import "fmt"

func conversion() {
	// Kode Program Konversi Tipe Data (1)
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// # Hati hati pada saat mengkonversi nilai integer, apabila nilai integer tidak sanggup maka dia akan mereturn data terbesar dari integer tersebut (integer overflow)

	// Kode Program Konversi Tipe Data (2)
	var name = "Luqman Zakariya"
	var e byte = name[0] // tipe data byte
	var eString = string(e)

	fmt.Println(name)
	fmt.Println((eString))

}
