package main

/* === File Reference ===
hook_test.go
*/

/* === Hook ===
- Hook adalah sebuah Struct yang bisa kita tambahkan ke Logger sebagai callback yang akan dieksekusi
	ketika terdapat kejadian log untuk level tertentu
- Contohnya misal, ketika ada log error, dan kita ingin mengirim notifikasi via chat ke programmer dan
	lain-lain
- Kita bisa menambah Hook ke Logger dengan menggunakan function logger.AddHook()
- Dan kita juga bisa menambahkan lebih dari satu Hook ke Logger
*/
