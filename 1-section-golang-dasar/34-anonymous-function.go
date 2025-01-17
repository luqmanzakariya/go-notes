package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func main() {
// 	// Anynymous Function
// 	// Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
// 	// Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
// 	// Hal tersebut dinamakan anonymous function, atau function tanpa nama

// 	blacklist := func(name string) bool {
// 		return name == "Anjing"
// 	}
// 	registerUser("Luqman", blacklist)

// 	// Or

// 	registerUser("Anjing", func(name string) bool {
// 		return name == "Anjing"
// 	})
// }
