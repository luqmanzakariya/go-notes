package main

/* === Aturan File Test ===
- Go-Lang memiliki aturan cara membuat file khusus untuk unit test
- Untuk membuat file unit test, kita harus menggunakan akhiran _test
- Jadi kita misal kita membuat file hello_world.go, artinya untuk membuat unit testnya, kita harus
	membuat file hello_world_test.go
*/

/* === Aturan Function Unit Test ===
- Selain aturan nama file, di Go-Lang juga sudah diatur untuk nama function unit test
- Nama function untuk unit test harus diawali dengan nama Test
- Misal jika kita ingin mengetest function HelloWorld, maka kita akan membuat function unit test
	dengan nama TestHelloWorld
- Tak ada aturan untuk nama belakang function unit test harus sama dengan nama function yang akan
	di test, yang penting adalah harus diawalin dengan kata Test
- Selanjutnya harus memiliki parameter (t *testing.T) dan tidak mengembalikan return value
*/

/* === Menjalankan Unit Test ===
- Untuk menjalankan unit test kita bisa menggunakan perintah (akan mencari semua file golang yang
	berakhiran _test):
	go test
- Jika kita ingin lihat lebih detail function test apa saja yang sudah di running, kita bisa guakan
	perintah:
	go test -v
- Dan jika kita ingin memilih function unit test mana yang ingin di running, kita bisa gunakan perintah:
	go test -v -run TestNamaFunction
*/

/* === Menjalankan Semua Unit Test ===
- Jika kita ingin menjalankan semua unit test dari top folder module nya, kita bisa gunakan perintah:
	go test ./...
*/