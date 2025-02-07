package main

/* === File Reference ===
formatter_test.go
*/

/* === Formatter ===
- Saat Logger mengirim data ke Output, log yang kita kirim akan diformat menggunakan object Formatter
- Logrus secara default memiliki dua formatter:
- TextFormatter, yang secara default digunakan
- JSONFormatter, yang bisa digunakan untuk memformat pesan log menjadi data JSON
- Untuk mengubah formatter, kita bisa gunakan function logger.SetFormatter()
*/
