package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/* === Header ===
- Selain Query Parameter, dalam HTTP, ada juga yang bernama Header
- Header adalah informasi tambahan yang biasa dikirim dari client ke server atau sebaliknya
- Jadi dalam Header, tidak hanya ada pada HTTP Request, pada HTTP Response pun kita bisa menambahkan
	informasi header
- Saat kita menggunakan browser, biasanya secara otomatis header akan ditambahkan oleh browser,
	seperti informasi browser, jenis tipe content yang dikirim dan diterima oleh browser, dan lain-lain
*/

/* === Request Header ===
- Untuk menangkap request header yang dikirim oleh client, kita bisa mengambilnya di Request.Header
- Header mirip seperti Query Parameter, isinya adalah map[string][]string
- Berbeda dengan Query Parameter yang case sensitive, secara spesifikasi, Header key tidaklah case sensitive
*/

/* === Response Header ===
- Sedangkan jika kita ingin menambahkan header pada response, kita bisa menggunakan function 
	ResponseWriter.Header()
*/

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")
	fmt.Fprintln(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json") // Header case insensitive tidak seperty query

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Programmer Zaman Now")
	fmt.Fprintf(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	request.Header.Add("Content-Type", "application/json") // Header case insensitive tidak seperty query

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	
	fmt.Println(response.Header.Get("x-powered-by")) // Cek value header x-powered-by
}