package main

/* === Embed Package ===
- Sejak Golang versi 1.16, terdapat package baru dengan nama embed
- Package embed adalah fitur baru untuk mempermudah membaca isi file pada saat compile time secara
	otomatis dimasukkan isi file nya ke dalam variable
- https://golang.org/pkg/embed/
*/

// Untuk mengikuti course ini pastikan versio go lang ada di versi 1.16 keatas

/* === Cara Embed File ===
- Untuk melakukan embed file ke variable, kita bisa mengimport package embed terlebih dahulu
- Selanjutnya kita bisa tambahkan komentar //go:embed diikuti dengan nama file nya, diatas variable yang kita
	tuju
- Variable yang dituju tersebut nanti secara otomatis akan berisi konten file yang kita inginkan secara
	otomatis ketika kode golang di compile
- Variable yang dituju tidak bisa disimpan di dalam function
*/