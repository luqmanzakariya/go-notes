package main

// import (
// 	"fmt"
// 	"time"
// )

/* === Package time ===
- Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
*/

/* # Contoh package time
- time.Now() : Untuk mendapatkan waktu saat ini
- time.Date(): Untuk membuat waktu
- time.Parse(layout, string) : Untuk memparsing waktu dari string
*/

/* === Duration ===
- Saat menggunakan tipe data waktu, kadang kita butuh data berupa durasi
- Package tipe memiliki type Duration, yang sebenearnya adalah alias untuk int64
- Namun terdapat banyak method yang bisa kita gunakan untuk memanipulasi Duration
*/

// func main() {
// 	now := time.Now() // Membuat waktu saat ini berdasarkan komputer yang digunakan
// 	fmt.Println(now.Local()) // Untuk mendapatkan waktu, (waktu sesuai komputer yang digunakan)

// 	var utc time.Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
// 	fmt.Println(utc) // Waktu utc
// 	fmt.Println(utc.Local()) // waktu utc dalam lokal

// 	// Format Waktu 1 (formatter harus sesuai dengan waktu yang ada didokumentasi golang)
// 	formatter := "2006-01-02 15:04:05"
// 	value := "2020-10-10 10:10:10"
// 	valueTime, err := time.Parse(formatter, value)
// 	if err != nil {
// 		fmt.Println("Error", err.Error())
// 	} else {
// 		fmt.Println(valueTime)
// 	}

// 	// Format Waktu 2 (formatter harus sesuai dengan waktu yang ada didokumentasi golang)
// 	parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
// 	fmt.Println(parse)

// 	// Get hour detail
// 	fmt.Println(valueTime.Year())
// 	fmt.Println(valueTime.Month())
// 	fmt.Println(valueTime.Day())
// 	fmt.Println(valueTime.Hour())



// 	// # DURATION
// 	var duration1 time.Duration = time.Second * 100
// 	var duration2 time.Duration = time.Microsecond * 10
// 	var duration3 time.Duration = duration1-duration2

// 	fmt.Println(duration3)
// 	fmt.Printf("duration1: %d \n", duration3)
// }
