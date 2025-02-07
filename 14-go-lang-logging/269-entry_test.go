package main

/* === File Reference ===
entry_test.go
*/

/* === Entry ===
- Entry adalah sebuah Struct representasi dari log yang kita kirim menggunakan Logrus Logger
- Setiap log yang kita kirim, maka akan dibuatkan object Entry
- Contohnya ketika kita membuat formatter sendiri, maka parameter yang kita dapat untuk melakukan
	formatting bukanlah string message, melainkan object Entry
- https://github.com/sirupses/logrus/master/entry.go
- Untuk membuat entry, kita bisa menggunakan function logrus.NewEntry()
*/
