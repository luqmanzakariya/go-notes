package main

import (
	"context"
	"fmt"
	"testing"
)

/* === Context With Value ===
- Pada saat awal membuat context, context tidak memiliki value
- Kita bisa menambah sebuah value dengan data Pair (key-value) ke dalam context
- Saat kita menambah value ke context, secara otomatis akan tercipta child context baru, artinya original
	context nya tidak akan berubah sama sekali
- Untuk membuat menambahkan value ke context, kita bisa menggunakan function
	context.WithValue(parent, key, value)
*/

// Kode: Context With Value
func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	// Check the value in terminal using: go test -v -run=TestContextWithValue
	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// Context Get Value: Context akan terus mencari value dari child sampai ke parentnya
	fmt.Println(contextF.Value("f")) // Dapet (ada di child)
	fmt.Println(contextF.Value("c")) // Dapet (karena value ada di parent)
	fmt.Println(contextF.Value("g")) // <nil> (karena valuenya tidak ada di child dan di parent)
	fmt.Println(contextA.Value("b")) // <nil> (valuenya kosong karena context hanya nanya ke parent, ke child 
																	 //        tidak akan ditanya)
}

