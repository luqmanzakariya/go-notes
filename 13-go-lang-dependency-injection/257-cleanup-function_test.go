package main

/* === Cleanup Function ===
- Jika Provider membuat object yang membutuhkan proses cleanup (pembersihan) setelah object dibuat, maka pada provider
	kita bisa mengembalikan closure
- Closure secara otomatis akan dipanggil dalam proses cleanup oleh Google Wire
*/

/* === File Reference ===
- belajar-golang-dependency-injection/simple/file.go
- belajar-golang-dependency-injection/simple/connection.go
- belajar-golang-dependency-injection/simple/injector.go
- belajar-golang-dependency-injection/test/file_test.go
*/