package main

// import "fmt"

// === Nil ===
// Biasanya di dalam bahasa pemorgraman lain, object yang belum diinisialisasi maka secara otomatis nilainya adalah null atau nil
// Berbeda dengan Go-Lang, di Go-Lang saat kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default valuenya
// Namun di Go-Lang ada data nil, yaitu data kosong
// Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer, dan channel

// Kode Program Nil(1)
func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

// func main() {
// 	var person map[string]string = nil
// 	fmt.Println(person)

// 	person2 := NewMap(("Luqman"))
// 	fmt.Println(person2)

// 	var person3 map[string]string
// 	if person3["name"] == "" {
// 		fmt.Println("Data kosong")
// 	} else {
// 		fmt.Print("Halo", person3)
// 	}
// }
