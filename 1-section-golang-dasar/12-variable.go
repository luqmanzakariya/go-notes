/*

*/

package main

import "fmt"

func variableFunc() {
	/* Variable Declaration */
	var person string

	person = "Luqman Zakariya"
	fmt.Println("person:", person)

	person = "Yolanda Puspita Sari Nasution"
	fmt.Println("person:", person)


	/* Variable Declaration Shortcut */
	var fullName = "Muhammad Luqman Zakariya"
	fmt.Println("fullName:", fullName)

	/* Variable Declaration Shortcut without var */
	// Syarat: Harus langsung dimasukkan datanya
	name := "Luqman Zakariya"
	fmt.Println("Name:", name)

	// Deklarasi multiple variable
	var (
		firstName = "Luqman"
		lastName = "Zakariya"
	)
	fmt.Println("Multiple Variable:", firstName, lastName)
}

// Variable yang telah dideklarasi tidak dapat diubah tipe datanya
