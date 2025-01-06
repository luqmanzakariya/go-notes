package main

import (
	"fmt"
)

// === Struct Method ===
// Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
// Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
// Method adalah function

type Customer2 struct {
	Name, Address string
	Age           int
}

func (customer Customer2) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is:", customer.Name)
}

func (a Customer2) sayBuu(name string) {
	fmt.Println(a.Name, "is the best", "what do you think,", name, "?")
}

// func main() {
// 	person := Customer2{
// 		Name:    "Luqman",
// 		Address: "Indonesia",
// 		Age:     30,
// 	}

// 	person.sayHi("Ola")
// 	person.sayBuu("Ola")
// }
