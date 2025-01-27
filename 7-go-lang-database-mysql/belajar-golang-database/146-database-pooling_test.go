package belajar_golang_database

/* === Membuat Koneksi ke Database ===
- sq.DB di Golang sebenarnya bukanlah sebuah koneksi ke database
- Melainkan sebuah pool ke database, atau dikenal dengan konsep Database Pooling
- Di dalam sql.DB, Golang melakukan management koneksi ke database secara otomatis. Hal ini menjadikan kita
	tidak perlu melakukan management koneksi database secara manual
- Dengan kemampuan database pooling ini, kita bisa menentukan jumlah minimal dan maksimal koneksi yang dibuat
	oleh Golang, sehingga tidak membanjiri koneksi ke database, karena biasanya ada batas maksimal koneksi
	yang bisa ditangani oleh database yang kita gunakan
*/

/* === Pengaturan Database Pooling ===
Method															Keterangan
(DB) SetMaxIdleConns(number)				Pengaturan berapa jumlah koneksi minimal yang dibuat
(DB) SetMaxOpenConns(number)				Pengaturan berapa jumlah koneksi maksimal yang dibuat
(DB) SetConnMaxIdleTime(duration)		Pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus
(DB) SetConnMaxLifetime(duration)		Pengaturan berapa lama koneksi boleh digunakan
*/

// Code: see file database.go