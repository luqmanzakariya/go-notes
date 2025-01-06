package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

/* === Package io ===
- Package bufio atau singkatan dari buffered io
- package ini digunakan untuk membuat data IO seperti Reader dan Writer
*/

// func main() {
// 	// # Reader
// 	input := strings.NewReader("this is long string \nwith new line\n")

// 	reader := bufio.NewReader(input)

// 	for {
// 		line, _, err := reader.ReadLine()
// 		if err == io.EOF {
// 			break
// 		}

// 		fmt.Println(string(line))
// 	}

// 	// # Reader
// 	writer := bufio.NewWriter(os.Stdout)
// 	writer.WriteString("hello world \n")
// 	writer.WriteString("selamat belajar")
// 	writer.Flush()
// }
