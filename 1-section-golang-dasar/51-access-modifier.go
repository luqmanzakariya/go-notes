package main

/* === Access Modifier ===
- Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access
  modifier terhadap suatu function atau variable
- Di Go-Lang, untuk menentukan access modifier, cukup nama function atau variable
- Jika nama nya diawali dengan huruf besar, maka artinya bisa diakses dari package lain, jika
  dimulai dengan hurup kecil, artinya tidak bisa diakses dari package lain
*/

// import (
// 	"learn-go/helper" // "learn-go" >> diambil dari nama module di go.mod
// 	"fmt"
// )

// func main () {
// 	appType := helper.Application // Bisa diakses
// 	// appVersion := helper.version // Bisa diakses
// 	fmt.Println("Application type", appType)
// 	// result := helper.SayHello("Luqman") // cara importnya: <nama package/folder>/<nama function>
// 	// fmt.Println(result)
// }