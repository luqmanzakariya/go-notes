package main

// import "fmt"

/* === Pointer di Method ===
- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
*/

type Man struct {
	Name string
}

// # Tanpa Pointer
// func (man Man) Married() {
// 	man.Name = "Mr. " + man.Name
// }

// func main(){
// 	luqman := Man{"Luqman"}
// 	luqman.Married()

// 	fmt.Println(luqman.Name)
// }

// # Dengan Pointer
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

// func main(){
// 	luqman := Man{"Luqman"}
// 	luqman.Married()

// 	fmt.Println(luqman.Name)
// }