package main

import (
	"embed"
	"fmt"
	"testing"
)

/* === Embed Multiple Files ===
- Kadang ada kebutuhan kita ingin melakukan embed beberapa file sekaligus
- Hal ini juga bisa dilakukan menggunakan embed package
- Kita bisa menambahkan komentar //go:embed lebih dari satu baris
- Selain itu variable nya bisa kita gunakan tipe data embed.FS
*/

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T){
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a)) // Karena tipe datanya adalah []bytes, perlu dikonversi dulu ke string
	
	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b)) // Karena tipe datanya adalah []bytes, perlu dikonversi dulu ke string

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c)) // Karena tipe datanya adalah []bytes, perlu dikonversi dulu ke string
}