package main

/* === Benchmark Function ===
- Mirip seperti unit test, untuk benchmark pun, di Go-Lang sudah ditentukan nama functionnya,
	harus diawali dengan kata Benchmark, misal BenchmarkHelloWorld, BenchmarkXxx
- Selain itu, harus memiliki parameter (b *testing.b)
- Dan tidak boleh mengembalikan return value
- Untuk nama file benchmark, sama seperti unit test, diakhiri dengan _test, misal hello_world_test.go
*/

/* === Contoh Benchmark ===
- Ada di helper/hello_world_test.go dengan nama function BenchmarkHelloWorld
*/

/* === Menjalankan Benchmark ===
- Untuk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah sama seperti test,
	namun ditambahkan parameter bench:
	go test -v -bench=
- Jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah:
	go test -v -run=NotMatchUnitTest -bench=
- Kode diatas selain menjalankan benchmark, akan menjalankan unit test juga, jika kita hanya ingin
	menjalankan benchmark tertentu, kita bisa gunakan perintah:
	go test -v -run=NotMatchUnitTest -bench=BenchmarkTest
- Jika kita menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa
	gunakan perintah:
	go test -v -run -bench=. ./...
*/