package main

// import (
// 	"fmt"
// 	"path"
// 	"path/filepath"
// )

/* === Package path ===
- Package path digunakan untuk memanipulasi data path seperti path di URL atau path di File system
- Secara default Package path menggunakan slash sebagai karakter path nya, oleh karena itu cocok
	untuk data URL
- Namun jika ingin menggunakan untuk memanipulasi path di File System, karena Windows menggunakan
	backslash, maka khusus untuk File System, perlu menggunakan package path/filepath
*/

// func main() {
// 	// # Package path
// 	fmt.Println(path.Dir("hello/world.go")) // Mendapatkan directory
// 	fmt.Println(path.Base("hello/world.go")) // Mendapatkan base
// 	fmt.Println(path.Ext("hello/world.go")) // Mendapatkan Extension
// 	fmt.Println(path.Join("hello", "world", "main.go")) // Menggabungkan path

// 	// # Package path/filepath
// 	fmt.Println("")
// 	fmt.Println(filepath.Dir("hello/world.go"))
// 	fmt.Println(filepath.Base("hello/world.go"))
// 	fmt.Println(filepath.Ext("hello/world.go"))
// 	fmt.Println(filepath.IsAbs("hello/world.go"))
// 	fmt.Println(filepath.IsLocal("hello/world.go"))
// 	fmt.Println(filepath.Join("hello", "world", "main.go"))
// }
