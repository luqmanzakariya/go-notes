package main

/* === Library Dependency Injection ===
- Banyak sekali library Dependency Injection yang bisa kita gunakan di Go-Lang, misalnya:
- https://github.com/google/wire
- https://github.com/uber-go/fx
- https://github.com/golobby/container
- Dan lain-lain
*/

/* === Google Wire ===
- Pada kelas ini, kita akan menggunakan Google Wire sebagai Dependency Injection library nya
- Salah satu kenapa Google Wire menjadi piliha, karena saat ini Google Wire adalah library paling
	populer untuk melakukan Dependency Injection di Go-Lang
- Selain itu, Google Wire merupakan library Dependency Injection yang berbasis compile, artinya
	kodenya akan di generate, bukan menggunakan reflection
- Hal ini membuat Google Wire menjadi cepat, karena hasil kompilasi nya adalah kode yang sudah di
	generate melakukan Dependency Injection, tanpa perlu menggunakan reflection lagi
*/
