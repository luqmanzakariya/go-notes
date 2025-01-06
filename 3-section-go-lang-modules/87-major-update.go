package main

/* === Major Upgrade ===
- Major upgrade biasanya terjadi dikarenakan ada perubahan pada isi kode program kita, sehingga membuatnya
	tidak backward compatible
- Sebaiknya hal ini sebisa mungkin dihindari
- Namun jika tidak bisa dihindary, strategy terbaik adalah merubah nama module
*/

/* === Langkah merubah nama module (souce github)
1. pada file go.mod dapat dilihat module seperti dibawah
module github.com/luqmanzakariya/go-modules-example

2. Ganti atau tambahkan module dengan nama yang baru, contoh ditambahkan v2
module github.com/luqmanzakariya/go-modules-example/v2

3. Ubah tag major version:
git tag v2.0.0

4. Push tag terbaru 
git push origin v2.0.0

5. Pada aplikasi (folder app), di file go.mod, delete module related tersebut
require github.com/luqmanzakariya/go-modules-example v1.2.0

6. Download kembali versi v2 package tersebut
go get github.com/luqmanzakariya/go-modules-example/v2

# Catatan: Apabila didownload tanpa tambahan v2, maka akan terdownload versi terakhir yaitu v1.2.0. Karena
	package terbaru berbeda nama module
*/