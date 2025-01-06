package main

// import (
// 	"flag"
// 	"fmt"
// )

/* === Package Flag ===
- Package flag berisikan fungsionalitas untuk memparsing command line argument
*/

// func main() {
// 	var host *string = flag.String("host", "localhost", "database host") // Returnnya adalah sebuah pointer
// 	var port *int =  flag.Int("port", 0, "database port") // Returnnya adalah sebuah pointer
// 	username := flag.String("username", "root", "database username") // Returnnya adalah sebuah pointer
// 	password := flag.String("password", "root", "database password") // Returnnya adalah sebuah pointer

// 	flag.Parse() // untuk melakukan parsing perlu melakukan command berikut

// 	fmt.Println("Username", *username) // Karena returnnya pointer maka perlu asterisk operator 
// 	fmt.Println("Password", *password) // Karena returnnya pointer maka perlu asterisk operator 
// 	fmt.Println("Host", *host) // Karena returnnya pointer maka perlu asterisk operator 
// 	fmt.Println("Port", *port) // Karena returnnya pointer maka perlu asterisk operator 
// }

/* # Contoh command 1 (default/tanpa data input)
go run 63-package-flag.go

Username root
Password root
Host localhost
Port 0
*/

/* # Contoh command 2 (default/dengan data input)
go run 63-package-flag.go -username=luqman -password="rahasia banget" -host=199.99.99.99 -port=5000

Username luqman
Password rahasia banget
Host 199.99.99.99
Port 5000
*/