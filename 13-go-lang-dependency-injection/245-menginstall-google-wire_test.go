package main

/* === Menginstall Wire ===
- Google Wire membutuhkan aplikasi command line wire untuk melakukan auto generate kode Dependency
	injection ketika nanti membuat kode
- Program ini perlu kita install manual di komputer kita dengan perintah:
	go install github.com/google/wire/cmd/wire@latest
- Secara otomatis akan ada file binary di $GOPATH/bin/wire
- Agar aplikasi command line wire tersebut bisa diakses, jangan lupa masukkan ke $PATH sistem
	operasi kita, seperti yang pernah kita lakukan ketika belajar setting $PATH Go-Lang di kelas
	Go-Lang Dasar
*/

/* === Command untuk mengecek file GO PATH ===
- env | grep GOPATH (command di termial)
	/Users/luqmanzakariya/go (response di terminal)
- cd /Users/luqmanzakariya/go change directory (command di termial)
- cd bin (command di termial)
- ls -l (command di termial)
*/

/* === Check apakah google wire sudah terinstall ===
- wire help (command di terminal)
*/