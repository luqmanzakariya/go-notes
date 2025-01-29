package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

/* === Middleware ===
- Dalam pembuatan web, ada konsep yang bernama middleware atau filter atau interceptor
- Middleware adalah sebuah fitur dimana kita bisa menambahkan kode sebelum dan setelah sebuah handler
	di eksekusi
*/

/* === Middleware di Go-Lang web ===
- Sayangnya, di Go-Lang web tidak ada middleware
- Namun karena struktur handler yang baik menggunakan interface, kita bisa membuat middleware sendiri
	menggunakan handler
*/

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Before Execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Printf("After Execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request){
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(writer, "Hello Foo")
	})

	// Contoh Error Handler
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Panic Executed")
		panic("Ups")
	})

	LogMiddleware := LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: &LogMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/* === Error Handler ===
- Kadang middleware juga biasa digunakan untuk melakukan error handler
- Hal ini sehingga jika terjadi panic di Handler, kita bisa melakukan recover di middleware, dan
	mengubah panic tersebut menjadi error response
- Dengan ini, kita bisa menjaga aplikasi kita tidak berhenti berjalan
*/