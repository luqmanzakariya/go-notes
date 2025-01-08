package main

/* === Mengagalkan Unit Test ===
- Mengagalkan unit test menggunakan panic bukanlah hal yang bagus
- Go-Lang sendiri sudah menyediakan cara untuk mengagalkan unit test menggunakan Testing.T
- Terdapat function Fail(), FailNow(), Error(), dan Fatal() jika kita ingin menggagalkan unit test
*/

/* === t.Fail() dan t.FailNow() ===
- Terdapat dua function untuk menggagalkan unit test, yaitu Fail() dan FailNow(). Lantas apa bedanya?
- Fail() akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test. Namun diakhir ketika
	selesai, maka unit test tersebut dianggap gagal
- FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test
*/

/* === t.Error(args...) dan t.Fatal(args...) ===
- Selain Fail() dan FailNow(), ada juga Error() dan Fatal()
- Error() function lebih seperti melakukan log (print) error, namun setelah melakukan log error, dia
	akan secara otomatis memanggil funtion Fail(), sehingga mengakibatkan unit test dianggap gagal
- Namun karena hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai
- Fata() mirip dengan Error(), hanya saja, setelah melakukan log error, dia akan memanggil FailNow(),
	sehingga mengakibatkan eksekusi unit test berhenti
*/