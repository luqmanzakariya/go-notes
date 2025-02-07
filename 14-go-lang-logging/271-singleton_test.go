package main

/* === File Reference ===

*/

/* === Singleton ===
- Logrus sendiri memiliki singleton object untuk Logger, sehingga kita tidak perlu membuat 
	object Logger sendiri sebenarnya
- Namun artinya, jika kita ubah data Logger singleton tersebut, maka secara otomatis yang
	menggunakan logger tersebut akan berubah
- Secara default, logger singleton yang ada di logrus menggunakan TextFormatter dan Info Level
- Cara menggunakan Logger singleton ini, kita bisa langsung menggunakan package logrus nya saja
*/
