package main

/* === Package io ===
- IO atau singkatan dari Input Output, merupakan fitur di Goang yang digunakan sebagai standard untuk
	proses Input Output
- Di Golang, semua mekanisme input output pasti mengikuti standard package io
*/

/* # Reader
- Untuk membaca input, Golang menggunakan kontrak interface bernama Reader yang terdapat di package io

type Reader interface {
	Read(p []byte) (n int, err error)
}
*/

/* # Writer
- Untuk menulis ke output, Golang menggunakan kontrak interface bernama Writer yang terdapat di package io

type Writer interface {
	Write(p []byte) (n int, err error)
}
*/

/* # Implementasi IO
- Penggunaan dari IO sendiri di Golang terdapat dibanyak package, sebelumnya contohnya kita menggunakan
	CSV Reader dan CSV Writer
- Karena Package IO sebenarnya hanya kontrak untuk IO, untuk implementasinya kita harus lakukan sendiri
- Tapi untungnya, Golang juga menyediakan package untuk mengimplementasikan IO secara mudah, yaitu
	menggunakan package bufio
*/

