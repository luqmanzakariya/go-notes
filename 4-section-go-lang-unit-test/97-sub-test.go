package main

/* === Sub Test ===
- Go lang mendukung fitur pembuatan function unit test di dalam function unit test
- Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemrograman yang lainnya
- Untuk membuat sub test, kita bisa menggunakan function Run()
*/

/* === Menjalankan Hanya Sub Test ===
- Kita saudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah:
	go test -run TestNamaFunction
- Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah:
	go test -run TestNamaFunction/NamaSubTest
- Atau untuk semua test semua sub test di semua function, kita bisa gunakan perintah:
	go test -run /NamaSubTest
*/

/* === Contoh Command ===
	# Nge run sub test spesifik:
	go test -v -run TestSubTest/Eko
	# Nge run semua sub test yang ada kata "eko":
	go test -v -run /Eko
	
*/