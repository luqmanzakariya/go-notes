package belajar_golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

/* === Upload File ===
- Saat membuat web, selain menerima input data berupa form dan query param, kadang kita juga menerima input
	data berupa file dari client
- Go-Lang Web sudah memiliki fitur untuk management upload file
- Hal ini memudahkan kita ketika butuh membuat web yang menerima input file upload
*/

/* === MultiPart ===
- Saat kita ingin menerima upload file, kita perlu melakukan parsing terlebih dahulu menggunakan
	Request.ParseMultipartForm(size), atau kita bisa langsung ambil data file nya menggunakan
	Request.FormFile(name), di dalam nya secara otomatis melakukan parsing terlebih dahulu
- Hasilnya merupakan data-data yang terdapat pada package multipart, seperti multipart.File sebagai
	representasi file nya, dan multipart.FileHeader sebagai informasi file nya
*/

func UploadForm(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	request.FormFile("file")
	// request.ParseMultipartForm(32 << 20) // by default 32mb maksimal size (default memory)
	file, fileHeader, err := request.FormFile("file") // # Proses Mengambil File
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename) // # Proses Create Destinasi File
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file) // # Proses Simpan File ke Destinasi File
	if err != nil {
		panic(err)
	}

	name := request.PostFormValue("name") // Mengambil value form yang bukan upload file
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/sample.png
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Eko Kurniawan Khannedy")
	file, _ := writer.CreateFormFile("file", "CONTOHUPLOAD.png")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
