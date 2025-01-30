package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/* === Router ===
- Inti dari library HttpRouter adalah struct Router
- Router ini merupakan implementasi dari http.Handler, sehingga kita bisa dnegan mudah menambahkan
	ke dalam http.Server
- Untuk membuat Router, kita bisa menggunakan function httprouter.New(), yang akan mengembalikan
	Router pointer
*/

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello HttpRouter")
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}

/* === Http Method ===
- Router mirip dengan ServeMux, dimana kita bisa menambahkan route ke dalam Router
- Kelebihan dibandingkan dengan ServeMux adalah, pada Router, kita bisa menentukan HTTP Method yang ingin
	kita gunakan, misal GET, POST, PUT, dan lain-lain
- Cara menambahkan route ke dalam Router adalah gunakan function yang sama dengan HTTP Method nya, misal
	router.GET(), router.POST(), dan lain-lain
*/

/* === httprouter.Handle ===
- Saat kita menggunakan ServeMux, ketika menambah route, kita bisa menambahkan http.Handler
- Berbeda dengan Router, pada Router kita tidak menggunakan http.Handler lagi, melainkan menggunakan type
	httprouter.Handle
- Perbedaan dengan http.Handler adalah, pada httprouter.Handle, terdapat parameter ke tiga yaitu Params,
	yang akan kita bahas nanti di chapter sendiri
*/
