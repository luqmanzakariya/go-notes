package main

/* === Dependency Injection ===
- Setelah kita membuat Injector dan Provider, selanjutnya yang perlu kita lakukan adalah menggunakan aplikasi
	command line Google Wire untuk melakukan auto generate kode Dependency Injection
- Kita bisa menggunakan perintah ini untuk melakukan auto generate kode dependency injection wire gen namapackage
- Secara otomatis aplikasi Google Wire akan mencari kode Injector di package tersebut, lalu membuat file
	wire_gen.go yang isinya adalah kode otomatis dependency injection
*/

/* === Error sebelum menjalankan google wire VSCODE MAC ===
Apabila ketemu error ini dengan menggunakan vscode:
	No packages found for open file /Users/luqmanzakariya/Work/learn/learn-go/13-go-lang-dependency-injection/belajar-golang-dependency-injection/simple/injector.go.
	This file may be excluded due to its build tags; try adding "-tags=<build tag>" to your gopls "buildFlags" configuration
	See the documentation for more information on working with build tags:
	https://github.com/golang/tools/blob/master/gopls/doc/settings.md#buildflags.go list

Berikut langkah-langkahnya:
Solution 1: Add -tags=wireinject to VS Code (gopls settings)
	- Open Command Palette (Cmd + Shift + P on macOS).
	- Search for "Preferences: Open User Settings (JSON)".
	-	Add this to enable wireinject for gopls:
		"gopls": {
			"buildFlags": ["-tags=wireinject"]
		}
	- Save
*/

/* === Cara menjalankan command google wire ===
	- Apabila di root folder project (belajar-golang-dependency-injection)
		$ wire gen programmerzamannow/belajar-golang-restful-api/simple
	- Atau dapat langsung ke folder dimana file injector di letakkan
		$ cd belajar-golang-dependency-injection/simple
		$ wire
*/

/* 
	# Berikutnya untuk mengambil service kita tinggal memanggil package tersebut
	- Contoh file /test/simple_service_test.go
*/