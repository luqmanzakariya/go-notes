package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

/* === Hasil Embed di Compile ===
- Perlu diketahui, bahwa hasil embed yang dilakukan oleh package embed adalah permanent dan data file
	yang dibaca disimpan dalam binary golang nya
- Artinya bukan dilakukan secara realtime membaca file yang ada diluar
- Hal ini menjadikan jika binary file golang sudah di compile, kita tidak butuh lagi file external nya,
	dan bahkan jika diubah file external nya, isi variable nya tidak akan berubah lagi
*/

//go:embed version.txt
var Version string

//go:embed logo.png
var Logo []byte

//go:embed files/*.txt
var Path embed.FS

func main() {
	fmt.Println(Version)

	err := ioutil.WriteFile("logo_new.png", Logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntries, _ := Path.ReadDir("files")

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := Path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
