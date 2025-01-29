package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

/* === Redirect ===
- Saat kita membuat website, kadang kita butuh melakukan redirect
- Misal setelah selesai login, kita lakukan redirect ke halaman dashboard
- Redirect sendiri sebenarnya sudah standard di HTTP
	https://developer.mozilla.org/en-US/docs/Web/HTTP/Redirections
- Kita hanya perlu membuat response code 3xx dan menambah header Location
- Namun untungnya di Go-Lang, ada function yang bisa kita gunakan untuk mempermudah ini
*/

func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Redirect")
}

func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	// logic
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

func RedirectOut(writer http.ResponseWriter, request *http.Request) {
	// logic
	http.Redirect(writer, request, "https://luqmanzakariya.com", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)
	mux.HandleFunc("/redirect-out", RedirectOut)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
