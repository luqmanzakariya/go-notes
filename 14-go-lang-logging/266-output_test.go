package main

/* === File Reference ===
output_test.go
*/

/* === Output ===
- Secara default, output tujuan log yang kita kirim via Logrus adalah ke Console
- Kadang kita butuh mengubah output tujuan log, misal ke File atau Database
- Output Logger adalah io.Writer, jadi kita bisa menggunakan io.Writer dari Go-Lang atau bisa membuat
	sendiri mengikuti kontrak io.Writer
*/
