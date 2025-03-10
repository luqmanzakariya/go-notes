package belajar_golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

/* === FileServer ===
- Go-Lang memiliki sebuah fitur yang bernama FileServer
- Dengan ini, kita bisa membuat handler di Go-Lang web yang digunakan sebagai static file server
- Dengan menggunakan FileServer, kita tidak perlu manual me-load file lagi
- FileServer adalah handler, jadi bisa kita tambahkan ke dalam http.Server atau http.ServeMux
*/

/* === 404 Not Found ===
- Jika kita coba jalankan, saat kita membuka misal /static/index.js, maka akan dapat error 404 Not Found
- Kenapa ini terjadi?
- Hal ini dikarenakan FileServer akan membaca url, lalu mencari file berdasarkan url nya, jadi jika kita
	membuat /static/index.js, maka FileServer akan mencari ke file /resources/static/index.js
- Hal ini menyebabkan 404 Not Found karena memang file nya tidak bisa ditemukan
- Oleh karena itu, kita bisa menggunakan function http.StripPrefix() untuk menghapus prefix di url
*/

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	// mux.Handle("/static/", fileServer) // responnya akan 404 Not Found karena mencari ke /directory/static
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/* === Go-Lang Embed ===
- Di Go-Lang 1.16 terdapat fitur baru yang bernama Go-Lang embed
- Dalam Go-Lang embed kita bisa embed file ke dalam binary distribution file, hal ini mempermudah sehingga
	kita tidak perlu meng-copy static file lagi
- Go-Lang Embed juga memiliki fitur yang bernama embed.FS, fitur ini bisa diintegrasikan dengan FileServer
*/

/* === 404 Not Found ===
- Jika kita coba jalankan, dan coba buka /static/index.js, maka kita akan mendapatkan error 404 Not Found
- Kenapa ini terjadi? Hal ini karena di Go-Lang embed, nama folder ikut menjadi nama resource nya, misal
	resources/index.js jadi untuk mengaksesnya kita perlu gunakan url /static/resources/index.js
- Jika kita ingin langsung mengakes file index.js tanpa menggunakan resources, kita bisa menggunakan
	function fs.Sub() untuk mendapatkan sub directory
*/

//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T){
	directory, _ := fs.Sub(resources, "resources") // Pastikan nama foldernya harus sama
	// fileServer := http.FileServer(http.FS(resources)) // Perbedaan dengan menggunakan golang embed (tanpa fs.Sub)
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	// mux.Handle("/static/", fileServer) // responnya akan 404 Not Found karena mencari ke /directory/static tanpa golang embed
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // responnya akan 404 Not Found karena nama folder ikut menjadi nama resource nya di Go-Lang embed

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
