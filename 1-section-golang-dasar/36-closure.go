package main

import "fmt"

func closure() {
	// Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
	// Harap gunakan fitur closure ini dengan bjiak saat kita membuat aplikasi

	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}
	
	increment()
	fmt.Println("counter", counter)
}

// func main() {
// 	closure()
// }
