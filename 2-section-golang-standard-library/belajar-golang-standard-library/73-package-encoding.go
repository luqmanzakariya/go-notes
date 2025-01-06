package main

// import (
// 	"encoding/base64"
// 	"encoding/csv"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

/* === Package encoding ===
- Golang menyediakan package encoding untuk melakukan encode dan decode
- Golang menyediakan berbagai macam algoritma untuk encoding, contoh yang populer adalah base64,
  csv, dan json
*/

// func main() {
// 	// # Contoh base64
// 	fmt.Println("=== base64 ===")
// 	var encoded = base64.StdEncoding.EncodeToString([]byte("Eko Kurniawan Khannedy"))
// 	fmt.Println(encoded)

// 	var decoded, err = base64.StdEncoding.DecodeString(encoded)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println(string(decoded))
// 	}

// 	// # Contoh CSV Reader
// 	fmt.Println("")
// 	fmt.Println("=== CSV Reader ===")
// 	csvString := "eko,kurniawan,khannedy\n" +
// 		"budi,pratama,pratama\n" +
// 		"joko,morro,diah"
	
// 	reader := csv.NewReader(strings.NewReader(csvString))

// 	for {
// 		record, err := reader.Read()
// 		if err == io.EOF {
// 			break
// 		}

// 		fmt.Println(record)
// 	}

// 	// # Contoh CSV Writer
// 	fmt.Println("")
// 	fmt.Println("=== CSV Writer ===")
// 	writer := csv.NewWriter((os.Stdout))
// 	_ = writer.Write([]string{"eko", "kurniawan", "khannedy"})
// 	_ = writer.Write([]string{"budi", "pratama", "pratama"})
// 	_ = writer.Write([]string{"joko", "morro", "diah"})
// 	writer.Flush() // mengirim semua perubahan
// }
