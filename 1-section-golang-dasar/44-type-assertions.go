package main

// import "fmt"

// === Type Assertions ===
// Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
// Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong

func random() any {
	return "OK"
}

// func main() {
// 	var result any = random()
// 	// Apabila data tersebut tipe datanya berbeda maka dapat terjadi panic error seperti dibawah ini
// 	// var resultInt int = result.(int)
// 	// fmt.Println(resultInt)

// 	// Untuk mengantisipasi hal tersebut, kita dapat menggunakan kode program type assertions switch
// 	switch value := result.(type) {
// 	case string:
// 		fmt.Println("String", value)
// 	case int:
// 		fmt.Println("Int", value)
// 	default:
// 		fmt.Println("Unknown")
// 	}
// }
