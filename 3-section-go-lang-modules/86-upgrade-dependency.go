package main

/* === Upgrade Dependency ===
- Untuk upgrade dependency ke versi terbaru, kita bisa mengubah isi go.mod, lalu mengubah tag nya menjadi tag terbaru
- Untuk download versi terbaru, gunakan perintah: go get
*/

/* # Contoh command
1. buka folder go.mod (lihat folder app):
awal: require github.com/luqmanzakariya/go-modules-example v1.0.0

2. Ubah versi package
require github.com/luqmanzakariya/go-modules-example v1.2.0

3. Jalan kan di terminal
go get
*/

/*
	# Catatan: Apabila salah versi maka akan muncul error / buildnya akan gagal
	# Catatan: 	Apabila version sudah terdownload maka versi yang sama tidak akan terdownload lagi. Disarankan untuk 
							mengupdate versionnya walaupun mini seperti v1.2.1 karena di go apabila sudah terdownload maka akan
							terdapat cache pada version yang sama

*/