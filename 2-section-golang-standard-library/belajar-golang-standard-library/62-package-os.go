package main

/* === Package os ===
- Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
- Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa
  digunakan disemua sistem operasi)
*/

// import (
// 	"fmt"
// 	"os"
// )

// func main () {
// 	/* # 1. os.Args */
// 	args := os.Args

// 	for _, arg := range args {
// 		fmt.Println(arg)
// 	}

// 	/* # 2. os.Hostname() */
// 	hostname, err := os.Hostname()

// 	if err == nil {
// 		fmt.Println(hostname)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}

// }

/*
	contoh command: go run 62-package-os.go luqman zakariya
	/var/folders/qy/r615ng1d6jg8hc3_kyxy5cgh0000gn/T/go-build2905367799/b001/exe/62-package-os
	luqman
	zakariya
*/

/*
	contoh command: go run 62-package-os.go "luqman zakariya"
	/var/folders/qy/r615ng1d6jg8hc3_kyxy5cgh0000gn/T/go-build2853518763/b001/exe/62-package-os
	luqman zakariya
*/

/*
	contoh command: go run 62-package-os.go "luqman zakariya"
	Luqmans-MacBook-Pro.local
*/