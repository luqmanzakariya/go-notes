package main

// import (
// 	"container/ring"
// 	"fmt"
// 	"strconv"
// )

/* === Package container/ring ===
- Package container/ring adalah implementasi struktur data circular list
- Circular list adalah struktur data ring, dimana diakhir element akan kembali ke element awal (HEAD)
*/

// func main() {
// 	var data *ring.Ring = ring.New(5)

// 	// # 1 : Isi data manual
// 	// data.Value = "Value 1"

// 	// data = data.Next()
// 	// data.Value = "Value 2"

// 	// data = data.Next()
// 	// data.Value = "Value 3"
	
// 	// data = data.Next()
// 	// data.Value = "Value 4"

// 	// data = data.Next()
// 	// data.Value = "Value 5"
// 	// pada saat di print kenapa data 5 yang duluan karena saat ini sedang berada di data ke 5

// 	// # 2 : Isi data menggunakan looping
// 	for i := 0; i<data.Len(); i++ {
// 		data.Value = "Value " + strconv.Itoa(i)
// 		data = data.Next()
// 	}

// 	// Iterasi data menggunakan looping
// 	data.Do(func(value any){
// 		fmt.Println(value)
// 	})

// 	// # 3 Isi data dengan menggunakan looping
// 	// data := ring.New(5)

// 	// for i := 0; i < data.Len(); i++ {
// 	// 	data.Value = "Value" + strconv.FormatInt(int64(i), 10)
// 	// 	data = data.Next()
// 	// }

// 	// data.Do(func(value interface{}) {
// 	// 	fmt.Println(value)
// 	// })
// }
