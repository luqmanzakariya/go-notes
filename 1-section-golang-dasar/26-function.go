package main

import "fmt"

func functionWrapper(num int, name string) {
	fmt.Println(num+1, "Halo", name)
}

// func main() {
// 	// Secara default function main akan di run oleh go lang
// 	name := []string{"Luqman", "Ola", "Gasgas"}
	
// 	for i := 0; i < 3; i++ {
// 		functionWrapper(i, name[i])
// 	}
// }
