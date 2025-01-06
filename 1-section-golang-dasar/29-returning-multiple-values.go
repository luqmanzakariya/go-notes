package main

// import "fmt"

func getFullName() (string, string, string) {
	// Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multipe value
	// Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya di function

	return "Muhammad", "Luqman", "Zakariya"

	// Menghiraukan Return Value
	// Multiple return value wajib ditangkap semua value nya
	// Jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda _ (garis bawah)
}

// func main() {
// 	firstName, _, lastName := getFullName()
// 	fmt.Println(firstName)
// 	// fmt.Println(middleName)
// 	fmt.Println(lastName)
// }
