package belajar_golang_database

/* === Repository Pattern ===
- Dalam buku Domain-Driven Design, Eric Evans menjeleaskan bahwa "repository" is a mechanism for
	encapsulating storage, retrieval, and search behaviour, which emulates a collection of objects
- Pattern Repository ini biasanya digunakan sebagai jembatan antar business logic aplikasi kita dengan
	semua perintah SQL ke database
- Jadi semua perintah SQL akan ditulis di Repository, sedangkan business logic kode program kita hanya
	cukup menggunakan Repository tersebut
*/

/* === Entity / Model ===
- Dalam pemrograman berorientasi object, biasanya sebuah table di database akan selalu dibuat representasinya
	sebagai class Entity atau Model, namun di Golang, karena tidak mengenal class, jadi kita akan representasikan data
	dalam bentuk struct
- Ini bisa mempermudah ketika membuat kode program
- Misal ketika kita query ke Repository, dibanding mengembalikan array, alangkah baiknya Repository melakukan
	konversi terlebih dahulu ke struct Entity / Model, sehingga kita tinggal menggunakan objectnya saja
*/

// Untuk pengerjaan cek folder entity dan folder repository