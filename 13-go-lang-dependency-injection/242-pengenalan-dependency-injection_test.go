package main

/* === Dependency Injection ===
- Dalam pembuatan perangkat lunak, Dependency Injection merupakan sebuah teknik dimana sebuah object
	menerima object lain yang dibutuhkan (dependencies) ketika pembuatan object itu sendiri
- Biasanya object yang menerima dependencies disebut client, proses mengirim dependencies ke object
	tersebut biasa dibilang inject
- Dependency Injection sebenarnya sudah sering sekali kita lakukan, misal membuat object
	Controller yang membutuhkan dependencies object Service, atau membuat object service yang
	membutuhkan dependencies object Repository
*/

/* === Function Sebagai Constructor (1) ===
- Dalam bahasa pemrograman berorientasi object, ada istilah yang bernama Constructor, yaitu
	sebuah function yang digunakan ketika sebuah object dibuat
- Di Go-Lang, biasanya kita juga membuat sebuah function untuk membuat object, dan ini mirip
	seperti Constructor tugasnya, yaitu membuat object baru

	contoh:
	func NewCategoryController(categoryService service.CategoryService) CategoryController {
		return &CategoryControllerImpl {
			CategoryService : categoryService,
		}
	}
*/

/* === Function Sebagai Constructor (2) ===
- Biasanya kita akan membuat object dengan memanggil function Constructor tersebut, lalu
	mengirimkan dependencies yang dibutuhkan pada function Constructor tersebut
- Cara seperti ini mudah dilakukan ketika kode program aplikasi kita tidak terlalu besar
- Namun saat kode program aplikasi kita semakin besar, akan semakin sulit melakukan hal ini,
	terutama kita harus tahu urutan object mana yang harus dibuat terlebih dahulu
- Oleh karena ini, proses Dependency Injection sebenarnya bisa kita permudah dengan memanfaatkan
	library
*/

/* === Kode: Manual Dependency Injection ===
db := app.NewDB()
validate := validator.New()
categoryRepository := repository.NewCategoryRepository()
categoryService := service.NewCategoryService(categoryRepository, db, validate)
categoryController := controller.NewCategoryController(categoryService)
router := app.NewRouter(categoryController)

server := http.Server{
	Addr : "localhost:3000",
	Handler : middleware.NewAuthMiddleware(router)
}
*/

/* 
	Kode diatas terlihat baru satu object saja, bayangkan kalau aplikasi lebih besar dan banyak
	sekali objectnya. Jadi tujuan dependency injection agar mempermudah injection ke
	masing masing object atau library yang dibutuhkan
*/