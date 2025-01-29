package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/* === Form Post ===
- Saat kita belajar HTML, kita tahu bahwa saat kita membuat form, kita bisa submit datanya dengan method
	GET atau POST
- Jika menggunakan method GET, maka hasilnya semua data di form akan menjadi query parameter
- Sedangkan jika menggunakan POST, maka semua data di form akan dikirim via body HTTP request
- Di Go-Lang, untuk mengambil data Form Post sangatlah mudah
*/

/* === Request.PostForm ===
- Semua data form post yang dikirim dari client, secara otomatis akan disimpan dalam attribute Request.PostForm
- Namun sebelum kita bisa mengambil data di attribute PostForm, kita wajib memanggil method Request.ParseForm()
	terlebih dahulu, method ini digunakan untuk melakukan parsing data body apakah bisa di parsing menjadi
	form data atau tidak, jika tidak bisa di parsing maka akan menyebabkan error
*/

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	// request.PostFormValue("first_name") // Parsing otomatis tanpa menggunakan ParseForm()

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Eko&last_name=Khannedy")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
