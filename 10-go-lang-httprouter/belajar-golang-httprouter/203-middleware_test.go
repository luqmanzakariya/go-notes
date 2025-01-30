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

/* === Middleware ===
- HttpRouter hanyalah library untuk http router saja, tidak ada fitur selain router
- Dan karena Router merupakan implementasi dari http.Handler, jadi untuk middleware, kita bisa membuat sendiri,
	seperti yang sudah kita bahas pada course Go-Lang Web
*/

type LogMiddleware struct {
	http.Handler
}

func (middleware LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Receive request")
	middleware.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, response *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Middleware")
	})

	middleware := LogMiddleware{router}

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Middleware", string(body))
}
