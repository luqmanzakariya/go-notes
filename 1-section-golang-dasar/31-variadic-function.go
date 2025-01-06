package main

// import "fmt"

func sumAll(numbers ...int) int {
	// Parameter yang berada di posisi terakhi, memiliki kemampuan dijadikan sebuah varargs
	// Varargs artinya data bisa menerima lebih dari satu input, atau anggap saja semacam array
	// Apa bedanya dengan parameter biasa dengan tipe data array?
	// - Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
	// - Jika parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma

	// Kode Program Variadic Function
	total := 0
	for _, number := range numbers {
		total += number
	}

	// Argument cuma bisa ditaruh dipaling belakang (..int)
	return total
}

// func main() {
// 	total := sumAll(10, 10, 10, 10, 10)
// 	fmt.Println(total)

// 	// Dengan menggunakan variadic, parameter di function jadi tidak wajib (parameternya bisa kosong)
// 	total2 := sumAll()
// 	fmt.Println(total2)

// 	// === Slice Parameter ===
// 	// Kita bisa menjadikan slice sebagai vararg parameter

// 	// Kode Program Slice Parameter
// 	numbers := []int{5, 5, 5, 5, 5}
// 	total3 := sumAll(numbers...)
// 	fmt.Println(total3)
// }
