package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/* === Template Action ===
- Go-Lang template mendukung perintah action, seperti percabangan, perulangan, dan lain-lain
*/

/* === If Else ===
- {{if .Value}} T1 {{end}}, jika Value tidak kosong, maka T1 akan dieksekusi, jika kosong, tidak ada yang
	di eksekusi
- {{if .Value}} T1 {{else}} T2 {{end}}, jika value tidak kosong, maka T1 akan dieksekusi, jika kosong T2
	yang akan dieksekusi
- {{if .Value1}} T1 {{elseif.value2}} T2 {{else}}T3{{end}}, jika value1 tidak kosong, maka T1 akan dieksekusi,
	jika value2 tidak kosong, maka T2 akan dieksekusi, jika tidak semuanya, maka T3 akan dieksekusi
*/

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Eko",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/* === Operator Perbandingan ===
Go-Lang template juga mendukung operator perbandingan, ini cocok ketika butuh melakukan perbandingan,
ini cocok ketika butuh melakukan perbandingan number di if statement, berikut adalah operator nya:
- eq		artinya arg1 == arg2 	// equals
- ne		artinya arg1 != arg2 	// not equals
- lt		artinya arg1 < arg2		// less than
- le		artinya arg1 <= arg2	// less than equals
- gt		artinya arg1 > arg2		// greater than
-	ge		artinya arg1 >= arg2	// greater than equals
*/

/* === Kenapa Operatornya di Depan? ===
- Hal ini dikarenakan, sebenarnya operator perbandingan tersebut adalah sebuah function
- Jadi saat kita menggunakan {{eq First Second}}, sebenarnya dia akan memanggil function eq
	dengan paramter First dan Second: eq(First, Second)
*/

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FinalValue": 50,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/* === Range ===
- Range digunakan untuk melakukan iterasi data template
- Tidak ada perulangan biasa seperti menggunakan for di Go-Lang template
- Yang kita bisa lakukan adalah menggunakan range untuk mengiterasi tiap data array, slice, map atau
	channel
- {{range $index, $element := .Value}} T1 {{end}}, jika value memiliki data, maka T1 akan dieksekusi
	sebanyak element value, dan kita bisa menggunakan $index untuk mengakses index dan $element untuk
	mengakses element
- {{range $index, $element := .Value}} T1 {{else}} T2 {{end}}, sama seperti sebelumnya, namun jika
	value tidak memiliki element apapun, maka T2 yang akan dieksekusi
*/

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies": []string{
			"Game", "Read", "Code",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/* === With ===
- Kadang kita sering membuat nested struct
- Jika menggunakan template, kita bisa mengaksesnya menggunakan .Value.NestedValue
- Di template terdapat action with, yang bisa digunakan mengubah scope dot menjadi object yang kita mau
- {{with .Value}} T1 {{end}}, jika value tidak kosong, di T1 semua dot akan merefer ke value
- {{with .Value}} T1 {{else}} T2 {{end}}, sama seperti sebelumnya, namun jika value kosong, maka T2 yang
	akan dieksekusi
*/

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Title": "Template Action With",
		"Name": "Eko",
		"Address": map[string]interface{}{
			"Street": "Jalan Belum Ada",
			"City": "Jakarta",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

/* === Comment ===
- Template juga mendukung komentar
- Komentar secara otomatis akan hilang ketika template text di parsing
- Untuk membuat komentar sangat sederhana, kita bisa gunakan {{/* Contoh Komentar * /}}
*/

