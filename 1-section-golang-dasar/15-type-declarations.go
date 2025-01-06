package main

import "fmt"

func typeDeclaration () {
	type NoKTP string
	type Married bool
	

	var noKTPEko NoKTP = "12341234234"
	var marriedStatus Married = true
	fmt.Println(noKTPEko)
	fmt.Println(marriedStatus)

	// Mirip interface di typescript
}