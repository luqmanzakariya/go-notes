package main

// import (
// 	"fmt"
// 	"regexp"
// )

/* === Package regexp ===
- Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
- Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
*/

/* # Beberapa Function di Package regexp
- regexp.MustCompile(string): Membuat Regexp
- Regexp.MatchString(string) bool : Mengecek apakah Regexp match dengan string
- Regexp.FindAllString(string, max): Mencari string yang match dengan maximum jumlah hasil
*/

// func main() {
// 	var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

// 	fmt.Println(regex.MatchString("eko")) // true
// 	fmt.Println(regex.MatchString("edo")) // true
// 	fmt.Println(regex.MatchString("eKo")) // false

// 	fmt.Println(regex.FindAllString("eko edo egi ego elo eto eKo", 10))
// }
