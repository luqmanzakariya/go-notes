package main

/* === Binding Values ===
- Kadang ada kasus dimana kita ingin melakukan dependency injection terhadap value yang sudah ada,
	tanpa harus membuat provider terlebih dahulu
- Untuk kasus seperti ini, kita bisa langsung sebutkan value dari objectnya, tanpa menggunakan Provider
*/

/* === File Reference ===
- belajar-golang-dependency-injection/simple/injector.go
*/

/* === Interface Values ===
- Seperti di awal sudah dijelaskan, bahkan Google Wire akan melakukan dependency injection sesuai
	tipe data Provider nya
- Pada kasus jika kita ingin menggunakan value berupa Interface, maka kita perlu melakukan Interface
	Binding seperti yang sudah dibahas
- Atau ada cara yang lebih mudah, kita bisa binding value sekaligus menyebutkan interface yang digunakan
	oleh value tersebut
*/