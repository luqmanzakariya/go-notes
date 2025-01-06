package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// func main() {
// 	sayHelloWithFilter("Anjing", spamFilter)
// 	sayHelloWithFilter("Luqman", spamFilter)

// 	// Function Type Declaration
// 	// Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
// 	// Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter

// }
