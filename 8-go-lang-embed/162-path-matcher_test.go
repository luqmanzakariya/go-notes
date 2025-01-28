package main

import (
	"embed"
	"fmt"
	"testing"
)

/* === Path Matcher ===
- Selain manual satu per satu, kita bisa menggunakan path matcher untuk membaca multiple file yang kita
	inginkan
- Ini sangat cocok ketika misal kita punya pola jenis file yang kita inginkan untuk kita baca
- Caranya, kita perlu menggunakan path matcher seperti pada package function path.Match
*/

/* === func Match ===
Documentation: https://golang.org/pkg/path/#Match

func Match(pattern, name string) (matched bool, err error)

pattern:
	{ term }

term:
	'*'						matches any sequence of non-/ characters
	'?'						matches any single non-/ character
	'[' [ '^' ] 	{ character-range } ']'
								character class (must be non-empty)
	c							matches character c (c != '*', '?', '\\', '[')
	'\\' c				matches character c

character-range:
	c							matches character c (c != '\\', '-', ']')
	'\\' c				matches character c
	lo '-' hi			matches character c for lo <= c <= hi
*/

//go:embed files/*.txt
var path embed.FS
func TestPathMatcher(t *testing.T){
	dirEntries, _ := path.ReadDir("files")
	
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/"+ entry.Name())
			fmt.Println(string(file))
		}
	}
}