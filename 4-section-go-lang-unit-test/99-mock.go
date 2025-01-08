package main

/* === Mock ===
	- Mock adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil, dia akan 
		menghasilkan data yang sudah kita program diawal
	- Mock adalah salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object
		yang memang sulit untuk di testing
	- Misal kita ingin membuat unit test, namun ternyata ada kode program kita yang harus memanggil API Call ke
		third party service, dan belum tentu respon nya sesuai dengan apa yang kita mau
	- Pada kasus seperti ini, cocok sekali untuk menggunakan mock object
*/

/* === Testify Mock ===
- Untuk membuat mock object, tidak ada fitur bawaan Go-Lang, namun kita bisa menggunakan library testify yang
	sebelumnya kita gunakan untuk assertion
- Testify mendukung pembuatan mock object, sehingga cocok untuk kita gunakan ketika ingin membuat mock object
- Namun, perlu diperhatikan, jika desain kode program kita jelek, akan sulit untuk melakukan mocking, jadi
	pastikan kita melakukan pembuatan desain kode program kita dengan baik
- Mari kita buat contoh kasus
*/

/* === Aplikasi Query Ke Database ===
- Kita akan coba contoh kasus dengan membuat contoh aplikasi golang yang melakukan query ke database
- Dimana kita akan buat layer Service sebagai business logic, dan layer Repository sebagai jembatan ke database
- Agar kode kita mudah untuk di test, disarankan agar membuat kontrak berupa Interface
*/

/* === Langkah - Langkah ===
1. Pertama buat folder entity dan file category.go yang berisikan struct Category
2. Buat folder repository yang mana terdapat function yang akan berhubungan dengan database lalu buat kontraknya
3. Buat servicenya, idealnya membuat interfacenya tapi dalam kondisi ini kita langsung buat struct di file.
*/

/* === Langkah testing ===
1. Ke folder service
2. go test -v
*/