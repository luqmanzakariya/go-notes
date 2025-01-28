package main

import (
	_ "embed"
	"io/fs"
	"io/ioutil"
	"testing"
)

/* === Embed File ke []byte (dibaca slice of byte) ===
- Selain ke tipe data String, embed file juga bisa dilakukan ke variable tipe data []byte
- Ini cocok sekali jika kita ingin melakukan embed file dalam bentuk binary, seperti gambar
	dan lain-lain
*/

//go:embed logo.png
var logo []byte
func TestByte(t *testing.T){
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)

	if err != nil {
		panic(err)
	}
}