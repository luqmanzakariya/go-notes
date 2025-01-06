package main

// import (
// 	"fmt"
// 	"strconv"
// )

/* === Package strconv ===
- Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int34
- Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? Misal dari int ke string,
  atau sebaliknya
- Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
*/

/* # Contoh package strconv
- strconv.parseBool(string) : Mengubah string ke bool
- strconv.parseFloat(string) : Mengubah string ke float
- strconv.parseInt(string) : Mengubah string ke int64
- strconv.FormatBool(bool) : Mengubah bool ke string
- strconv.FormatFloat(float, ...) : Mengubah float65 ke string
- strconv.FormatInt(int, ...) : Mengubah int64 ke string
- strconv.Atoi(string) : Mengubah string to int
- strconv.Itoa(string) : Mengubah int to string
*/

// func main() {
// 	result, err := strconv.ParseBool("true")

// 	if err != nil {
// 		fmt.Println("Error", err.Error())
// 	} else {
// 		fmt.Println(result)
// 	}

// 	resultInt, errInt := strconv.Atoi("1000")

// 	if errInt != nil {
// 		fmt.Println("Error", errInt.Error())
// 	} else {
// 		fmt.Println(resultInt)
// 	}

// 	resultString := strconv.Itoa(-42)
// 	fmt.Println(resultString)

// 	binary := strconv.FormatInt(999, 2)
// 	fmt.Println(binary)
// }
