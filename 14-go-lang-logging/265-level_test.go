package main

/* === File Reference ===
level_test.go
*/

/* === Level ===
- Dalam Logging, hal yang paling penting adalah Level
- Level adalah penentuan prioritas atau jenis dari sebuah kejadian
- Level itu dimulai dari level terendah sampai level tertinggi
- Logrus mendukung banyak sekali level
*/

/* === Tabel Level di Logrus ===
Level			Function					Keterangan
Trace			logger.Trace()
Debug			logger.Debug()
Info			logger.Info()
Warn			logger.Warn()
Error			logger.Error()
Fatal			logger.Fatal()		Memanggil os.Exit setelah logging
Panic			logger.Panic()		Memanggil panic setelah logging
*/

/* === Logging Level ===
- Kenapa Trace dan Debug tidak keluar di console?
- Karena secara default, Logging Level yang digunakan adalah Info, artinya hanya prioritas Info keatas
	yang di log
- Untuk mengubah Logging Level, kita bisa gunakan logger.SetLevel()
*/
