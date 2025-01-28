package main

import (
	_ "embed"
	"fmt"
	"testing"
)

/* === Embed File ke String ===
- Embed file bisa kita lakukan ke variable dengan tipe data string
- Secara otomatis isi file akan dibaca sebagai text dan masukkan ke variable tersebut
*/

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}
