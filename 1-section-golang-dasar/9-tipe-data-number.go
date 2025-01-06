/*
	Tipe Data Number di Golang
	1. integer
	a. int8, int16, int32, int64
	b. uint8, uint16, uint32, uint64

	2. floating point
	a. float32, float64, complex64, complex128
	
	Tipe data Alias
	byte: uint8 (alias untuk uint8)
	rune: int32 (alias untuk int32)
	int: minimal int32
	uint: minimal uint32
*/

package main

import "fmt"

func tipeDataNumber () {
	fmt.Println("satu = ", 1)
	fmt.Println("dua = ", 2)
	fmt.Println("tiga koma lima = ", 3.5)
}