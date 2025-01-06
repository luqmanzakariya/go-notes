package main

/* === Membuat module ===
- Untuk membuat module baru, kita bisa menggunakan perintah: go mod init nama-module
- Go-Lang akan secara otomatis membuat file go.mod yang berisikan informasi nama module dan juga
	versi Go-Lang yang kita gunakan
*/

/* === Rilis Module ===
- Go-Lang terintegrasi baik dengan Git
- Untuk merilis module, kita hanya perlu membuat Tag di Git
*/

/* contoh command:
1. git commit -m "release 1"
2. git push origin master
3. git tag v1.0.0
4. git push origin v1.0.0
*/