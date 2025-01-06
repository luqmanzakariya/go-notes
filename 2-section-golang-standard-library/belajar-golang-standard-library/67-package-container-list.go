package main

// import (
// 	"container/list"
// 	"fmt"
// )

/* === Package container/list ===
- Package container/list adalah implementasi struktur data double linked list di Go-Lang

# note:
- Struktur data double linked list digunakan contohnhya untuk antrian atau tumpukan
*/

/* # Contoh package container/list
- list.New()
*/

// func main() {
// 	var data *list.List = list.New() // ini adalah pointer, pass by reference
// 	data.PushBack("Eko")
// 	data.PushBack("Kurniawan")
// 	data.PushBack("Khannedy")

// 	// data.Front() // Data yang paling pertama
// 	// data.Next() // Data selanjutnya
// 	// data.Value // Untuk ambil valuenya

// 	var head *list.Element = data.Front()
// 	fmt.Println(head.Value) // Eko

// 	next := head.Next()
// 	fmt.Println(next.Value) // Kurniawan

// 	next = next.Next()
// 	fmt.Println(next.Value) // Khannedy

// 	for e := data.Front(); e != nil; e = e.Next(){
// 		fmt.Println(e.Value)
// 	}
// }
