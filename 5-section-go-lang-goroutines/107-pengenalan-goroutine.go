package main

/* === Pengenalan Goroutine ===
- Goroutine adalah sebuah thread ringan yang dikelola oleh Go Runtime
- Ukuran Goroutine sangat kecil, sekitar 2kb, jauh lebih kecil dibandingkan Thread yang bisa sampai
	1mb atau 1000kb
- Namun tidak seperti thread yang berjalan parallel, goroutine berjalan secara concurrent
*/

/* === Cara Kerja Goroutine  ===
- Sebenarnya, Goroutine dijalankan oleh Go Scheduler dalam thread, dimana jumlah thread nya sebanyak
	GOMAXPROCS (biasanya sejumlah core CPU)
- Jadi sebenarnya tidak bisa dibilang Goroutine itu pengganti Thread, karena Goroutine sendiri berjalan
	diatas Thread
- Namun yang mempermudah kita adalah, kita tidak perlu melakukan manajemen Thread secara manual, semua
	sudah diatur oleh Go Scheduler
*/

/* === Cara Kerja Go Scheduler ===
Dalam Go-Scheduler, kita akan mengenal beberapa terminologi
- G : Goroutine
- M : Thread (Machine)
- P	: Processor
*/