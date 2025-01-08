package main

/* === Sub Benchmark ===
- Sama seperti testing.T, di testing.B juga kita bisa membuat sub benchmark menggunakan function Run()
*/

/* === Menjalankan Hanya Sub Benchmark ===
- Saat kita menjalankan benchmark function, maka semua sub benchmark akan berjalan
- Namun jika kita ingin menjalankan salah satu sub benchmark saja, kita bisa gunakan perintah:
	go test -v -bench=BenchmarkNama/NamaSub
*/