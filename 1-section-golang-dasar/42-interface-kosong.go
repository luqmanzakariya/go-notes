package main

// import "fmt"

// === interface kosong ===
// Go-Lang bukanlah bahasa pemrograman yang berorientasi objek
// Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
// Contoh di Java ada java.lang.Object
// Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
// Interface kosong adalah interface yang tidak memiliki deklarasi method satupun, hal ini membuat secara otomatis semua tipe data akan menjadi implementasi nya

// === Penggunaan Interface Kosong di Go-Lang ===
// Ada banyak contoh penggunaan interface kosong di Go-Lang, seperti:
// - fmt.Println(a ...interface{})
// - panic (v interface{})
// - recover() interface{}
// - dan lain-lain

// === Kode Program Interface Kosong ===
func Ups(i int, apapun interface{}) interface{} {
	if i == 1 {
		return "string"
	} else if i == 2 {
		return 2
	} else {
		return true
	}
}

// func main() {
// 	kosong1 := Ups(1, 2)

// 	// Karena returnnya adalah interface kosong, maka tipe data juga harus interface kosong
// 	var kosong2 interface{} = Ups(3, 2)

// 	fmt.Println(kosong1)
// 	fmt.Println(kosong2)
// }
