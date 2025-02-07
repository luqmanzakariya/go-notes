package main

/* === File Reference ===
field_test.go
*/

/* === Field ===
- Saat kita mengirim informasi log, kadang kita ingin menyisipkan sesuatu pada log tersebut
- Misal saja, kita ingin menyisipkan informasi siapa yang login di log nya
- Cara manual kita bisa menambahkan informasi di message nya, tapi Logrus menyediakan cara yang lebih
	baik, yaitu menggunakan fitur Field
- Dengan fitur Field, kita bisa menambahkan Field tambahan di informasi Log yang kita kirim
- Sebelum melakukan logging, kita bisa gunakan function logger.WithField() untuk menambahkan Field yang
	kita inginkan
*/

/* === Fields ===
- Atau, kita juga bisa langsung memasukkan beberapa Field dengan menggunakan Fields
- Fields adalah alias untuk map[string]interface{}
- Caranya kita bisa menggunakan function logger.WithFields()
*/
