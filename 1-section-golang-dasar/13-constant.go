package main

import "fmt"

func constantContainer() {
	const firstName string = "Luqman"
	const lastName = "Zakariya"
	fmt.Println("constant:", firstName, lastName)

	// Deklarasi constant harus langsung dikasih value (required / error jika tidak)
	// Constant tidak akan dikomplain apabila variable tidak digunakan

	// Deklarasi multiple constant
	const (
		depan = "Yolanda"
		belakang = "Sari"
	)
	fmt.Println("Multiple constant:", depan, belakang)
}
