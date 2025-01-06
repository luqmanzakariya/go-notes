package main

import (
	"bufio"
	// "fmt"
	"io"
	"os"
)

/* === File Management ===
- Di package os, terdapat File Management, namun sengaja ditunda pembahasannya, karena kita harus tau dulu
	tentang IO
- Saat kita membuat atau membaca file menggunakan Package os, struct File merupakan implementasi dari
	io.Reader dan io.Writer
- Oleh karena itu, kita bisa melakukan baca dan tulis terhadap File tersebut menggunakan Package io / bufio
*/

/* # Open File
- Untuk membuat/membaca File, kita bisa menggunakan os.OpenFile(nama, flag, permission)
- name berisikan nama file, bisa absolute atau relative/local
- flag merupakan penanda file, apakah untuk membaca, menulis, dan lain-lain
- permission, merupakan permission yang diperlukan ketika membuat file, bisa kita simulasikan disini:
  https://chmod-calculator.com/
*/

/* # File Flag di Package os
// Exactly one of O_RDONLY, O_wronly, or O_RDWR must be specified
O_RDONLY int = syscall.O_RDONLY // open the file read-only
O_WRONLY int = syscall.O_WRONLY // open the file write-only
O_RDWR   int = syscall.O_RDWR   // open the file read-write

// The remaining values may be or 'ed in to control behaviour
O_APPEND int = syscall.O_APPEND // append data to the file when writing
O_CREATE int = syscall.O_CREATE // create a new file if none exists.
O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist
O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O
O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened
*/

// Membuat File Baru
func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

// Membaca File
func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}		

		message += string(line) + "\n"
	}

	return message, nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

// func main() {
// 	// Example write
// 	createNewFile("sample.log", "this is sample log")

// 	// Example read
// 	result, _ := readFile("sample.log")
// 	fmt.Println(result)

// 	// Example add to file
// 	addToFile("sample.log", "this is add message")
// }
