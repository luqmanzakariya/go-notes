package main

import "fmt"

// Defer
// Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
// Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

func logging() {
	fmt.Println(("selesai memanggil function"))
}

func runApplication(value int) {
	// Apabila terjadi erro maka function defer tetap akan dieksekusi
	defer logging()
	result := 10 / value
	fmt.Println("result", result)
	fmt.Println("Run Application")
}

// Panic
// Panic function adalah function yang bisa kita gunakan untuk menghentikan program
// Panic function biasanya dipanggil ketika terjadi error pada saat program kita bejalan
// Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

func endApp() {
	// Kode Program Recover
	message := recover()
	
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}

	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()

	// # Setelah panic maka function dibawah panic tidak akan dieksekusi, bahkan jika terdapat function recover

	if error {
		panic("ALIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}

// Recover
// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
// Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
// Recover yang benar itu disimpan di defer function

// func main() {
// 	// Kode program defer
// 	runApplication(10)

// 	fmt.Println("====")

// 	// Kode Program Panic
// 	runApp(true)
// }
