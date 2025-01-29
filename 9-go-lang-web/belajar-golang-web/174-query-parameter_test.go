package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/* === Query Parameter ===
- Query parameter adalah salah satu fitur yang biasa kita gunakan ketika membuat web
- Query parameter biasanya digunakan untuk mengirim data dari client ke server
- Query parameter ditempatkan pada URL
- Untuk menambahkan query parameter, kita bisa menggunakan ?nama=value pada URL nya
*/

/* === url.URL ===
- Dalam parameter Request, terdapat attribute URL yang berisi data url.URL
- Dari data URL ini, kita bisa mengambil data query parameter yang dikirim dari client dengan menggunakan
	method Query() yang akan mengembalikan map
*/

/* === Multiple Query Parameter ===
- Dalam spesifikasi URL, kita bisa menambahkan lebih dari satu query parameter
- Ini cocok sekali jika kita memang ingin mengirim banyak data ke server, cukup tambahkan query parameter
	lainnya
- Untuk menambahkan query parameter, kita bisa gunakan tanda & lalu diikuti dengan query parameter
	berikutnya
*/

/* === Multiple Value Query Parameter ===
- Sebenarnya URL melakukan parsing query parameter dan menyimpannya dalam map[string][]string
- Artinya, dalam satu key query parameter, kita bisa memasukkan beberapa value
- Caranya kita bisa menambahkan query parameter dengan nama yang sama, namun value berbeda misal:
- name=Eko&name=Kurniawan
*/

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(writer, "Hello")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Eko", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleQueryParamter(writer http.ResponseWriter, request *http.Request){
	firstName := request.URL.Query().Get("first_name") // Kalau menggunakan Method Get maka akan mengambil data pada indekx ke o (first_name[0])
	lastName := request.URL.Query().Get("last_name") // Kalau menggunakan Method Get maka akan mengambil data pada indekx ke o (last_name[0])

	fmt.Fprintf(writer, "Helllo %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Eko&last_name=Khannedy", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParamter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParamterValues(writer http.ResponseWriter, request *http.Request){
	query := request.URL.Query()
	names := query["name"] // Dapetnya multiple kalau tidak menggunakan method GET
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Eko&name=Kurniawan&name=Khannedy", nil)
	recorder := httptest.NewRecorder()

	MultipleParamterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}