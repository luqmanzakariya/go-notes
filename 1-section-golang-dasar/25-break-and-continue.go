package main

import "fmt"

func breakContinue() {
	// Break digunakan untuk menghentikan seluruh perulangan
	// Continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya

	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("stop")
			break
		}
		fmt.Println("ini i ke", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("ini i ke", i)
	}

}