package main

/* === Binding Interface ===
- Dalam pembuatan aplikasi, sering kali kita biasanya menggunakan Interface sebagai kontrak struct
- Secara default, Google Wire akan menggunakan tipe data asli untuk melakukan dependency injection,
	jadi jika terdapat parameter berupa Interface, dan tidak ada Provider yang mengembalikan Interface
	tersebut, maka akan dianggap error
- Pada kasus ini, kita bisa memberi tahu ke Google Wire, jika kita ingin melakukan binding interface,
	yaitu memberi tahu untuk sebuah interface akan menggunakan provider dengan tipe apa
*/

/* === File Reference ===
- belajar-golang-dependency-injection/simple/hello.go
- belajar-golang-dependency-injection/simple/injector.go
*/