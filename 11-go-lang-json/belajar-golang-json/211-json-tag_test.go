package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/* === JSON Tag ===
- Secara default atribut yang terdapat di Struct dan JSON akan di mapping sesuai dengan nama atribut yang
	sama (case sensitive)
- Kadang ada stlye yang berbeda antara penamaan atribute di Struct dan di JSON, misal di JSON kita ingin
	menggunakan snake_case, tapi di Struct, kita ingin menggunakan PascalCase
- Untungnya, package json mendukung Tag Reflection
- Kita bisa menambahkan tag reflection dengan nama json, lalu diikuti dengan atribut yang kita inginkan
	ketika konversi dari atau ke JSON
*/

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Apple Mac Book Pro",
		ImageUrl: "http://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	// Untuk membaca json string go lang tidak peduli huruf besar atau kecilnya
	jsonString := `{"ID":"P0001","NAME":"Apple Mac Book Pro","IMAGE_URL":"http://example.com/image.png"}`
	jsonbyte := []byte(jsonString)

	product := &Product{}

	err := json.Unmarshal(jsonbyte, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
