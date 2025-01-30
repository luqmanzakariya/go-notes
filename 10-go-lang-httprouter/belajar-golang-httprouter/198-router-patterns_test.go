package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

/* === Router Pattern ===
- Sekarang kita sudah tahu bahwa dengan menggunakan Router, kita bisa menambah params di URL
- Sekarang pertanyaannya, bagaimana patter (pola) pembuatan parameter nya?
*/

/* === Named Parameter ===
- Named parameter adalah pola pembuatan parameter dengan menggunakan nama
- Setiap nama parameter harus diawali dengan :(titik dua), lalu diikuti dengan nama parameter
- Contoh, jika kita memiliki pattern seperti ini:

pattern								/user/:user
/user/eko							match
/user/you							match
/user/eko/profile			no match
/user/								no match
*/

/* === Catch All Parameter ===
- Selain named parameter, ada juga yang bernama catch all parameter, yaitu menangkap semua parameter
- Catch all parameter harus diawali dengan * (bintang), lalu diikuti dengan nama parameter
- Catch all parameter harus berada di posisi akhir URL

Pattern										/src/*filepath
/src/											no match
/src/somefile							match
/src/subdir/somefile			match
*/

func TestRouterPatternNamedParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id/items/:itemId", func(writer http.ResponseWriter, response *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		itemId := params.ByName("itemId")
		text := "Product " + id + " Item " + itemId
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1/items/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 Item 1", string(body))
}

func TestRouterPatternCatchAllParameter(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(writer http.ResponseWriter, response *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		text := "Image : " + image
		fmt.Fprint(writer, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image : /small/profile.png", string(body))
}