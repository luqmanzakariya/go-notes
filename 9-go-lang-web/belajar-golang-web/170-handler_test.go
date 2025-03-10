package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

/* === Handler ===
- Server hanya bertugas sebagai Web Server, sedangkan untuk menerima HTTP Request yang masuk ke Server,
	kita butuh yang namanya Handler
- Handler di Go-Lang di representasikan dalam interface, dimana dalam kontraknya terdapat sebuah
	function bernama ServeHTTP() yang digunakan sebagai function yang akan di eksekusi menerima HTTP
	request
*/

/* === HandlerFunc ===
- Salah satu implementasi dari interface Handler adalah HandlerFunc
- Kita bisa menggunakan HandlerFunc untuk membuat function handler HTTP
*/

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
