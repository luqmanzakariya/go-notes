package helper

/* === Package ===
- Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di Go-Lang
- Dengan menggunakan package, kita bisa merapikan kode program yang kita buat
- Package sendiri sebenarnya hanya direktori folder di sistem operasi kita
*/

func SayHello(name string) string { // Wajib menggunakan Huruf besar agar dapat diimport
	return "Hello " + name
}