package main

// import (
// 	"fmt"
// 	"strings"
// )

/* === Package Strings ===
- Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data String
- Ada banyak sekali function yang bisa kita gunakan
*/

/* # Contoh package strings
- strings.Trim(string, cutset): Memotong cutset di awal dan akhir string
- strings.ToLower(string): Membuat semua karakter string menjadi lower case
- strings.ToUpper(string): Membuat semua karakter string menjadi upper case
- strings.Split(string, separator): Memotong string berdasarkan separator
- strings.Contains(string, search): mengecek apakah string mengandung string lain
- strings.ReplaceAll(string, from, to): Mengubah semua string dari from ke to
*/

// func main () {
// 	fmt.Println(strings.Contains("Luqman Zakariya", "Luqman"))
// 	fmt.Println(strings.Split("Luqman Zakariya", " "))
// 	fmt.Println(strings.ToLower("Luqman Zakariya"))
// 	fmt.Println(strings.ToUpper("Luqman Zakariya"))
// 	fmt.Println(strings.Trim("         Luqman Zakariya          ", " "))
// 	fmt.Println(strings.ReplaceAll("Luqman Zakariya Luqman Zakariya", "Zakariya", "The Greatest"))
// }