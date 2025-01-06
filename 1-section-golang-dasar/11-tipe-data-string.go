/*
	String function
	len("String"): Menghitung jumlah karakter di String
	"string"[number]: Mengambil karakter pada posisi yang ditentukan
*/

package main

import "fmt"

func tipeDataString () {
	fmt.Println("Eko")
	fmt.Println("Eko Kurniawan")
	fmt.Println("Eko Kurniawan Khannedy")

	fmt.Println(len("Eko"))
	fmt.Println("Eko Kurniawan"[0]) // Responnya adalah bytenya dari "E"
	fmt.Println("Eko Kurniawan Khannedy"[1]) // Responnya adalah bytenya dari "k"
}