package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/* === Map ===
- Saat menggunakan JSON, kadang mungkin kita menemukan kasus data JSON nya dynamic
- Artinya atribute nya tidak menentu, bisa bertambah, bisa berkurang, dan tidak tetap
- Pada kasus seperti itu, menggunakan struct akan menyulitkan, karena pada Struct, kita harus
	menentukan semua atribut nya
- Untuk kasus seperti ini, kita bisa menggunakan tipe data map[string]interface{}
- Secara otomatis, atribut akan menjadi key di map, dan value menjadi value di map
- Namun karena value berupa interface{}, maka kita harus lakukan konversi secara manual jika ingin
	mengambil value nya
- Dan tipe data Map tidak mendukung JSON Tag lagi
*/

func TestMapDecode(t *testing.T){
	jsonString := `{"id": "p0001", "name": "Apple Mac Book Pro", "price": 20000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T){
	product := map[string]interface{}{
		"id": "P0001",
		"name": "Apple Mac Book Pro",
		"price": 20000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}