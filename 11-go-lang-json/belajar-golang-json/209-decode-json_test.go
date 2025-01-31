package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/* === Decode JSON ===
- Sekarang kita sudah tahu bagaimana caranya melakukan encode dari tipe data di Go-Lang ke JSON
- Namun bagaimana jika kebalikannya?
- Untuk melakukan konversi dari JSON ke tipe data di Go-Lang (Decode), kita bisa menggunakan fucntion
	json.Unmarshal(bypte[], interface{})
- Dimana byte[] adalah data JSON nya, sedangkan interface{} adalah tempat menyimpan hasil konversi, biasa
	berupa pointer
*/

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"Kurniawan","LastName":"Khannedy","Age":30,"Married":true}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}
