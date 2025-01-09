package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

/* === Membuat Channel ===
- Channel di Go-Lang direpresentasikan dengan tipe data chan
- Untuk membuat channel sangat mudah, kita bisa menggunakan make(), mirip ketika membuat map
- Namun saat pembuatan channel, kita harus tentukan tipe data apa saja yang bisa dimasukkan ke
	dalam channel tersebut
*/

/* === Mengirim dan Menerima Data dari Channel ===
- Seperti yang sudah dibahas sebelumnya, channel bisa digunakan untuk mengirim dan menerima data
- Untuk mengirim data, kita bisa gunakan kode: channel <-data
- Sedangkan untuk menerima data, bisa gunakan kode: data <-channel
- Jika selesai, jangan lupa untuk menutup channel menggunakan function close()
*/

/* === Perlu diperhatikan ketika menerima dan mengirim data dari Channel
- Apabila tidak ada data yang mengambil dari channel maka program akan terblock (panic: send on closed channel)
- Apabila mengambil data dari sebuah channel tapi kita tidak pernah mengirim data ke channel tersebut, maka akan 
	muncul error (fatal error: all goroutines are asleep - deadlock)
*/

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	
	// otomatis mengclose channel setelah function selesai (defer function)
	defer close(channel)

	// Mengirim data ke channel
	// channel <- "eko"

	// Menerima data dari channel
	// data := <- channel

	// Langsung mengirim channel sebagai parameter
	// fmt.Println(<- channel)

	// close channel
	// close(channel)

	go func(){
		time.Sleep(2 * time.Second)
		channel <- "Eko Kurniawan Khannedy" // Apabila tidak ada data yang mengambil dari channel maka program akan terblock (panic: send on closed channel)
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <- channel // Apabila mengambil data dari sebuah channel tapi kita tidak pernah mengirim data ke channel tersebut, maka akan muncul error (fatal error: all goroutines are asleep - deadlock)
	fmt.Println(data) 

	time.Sleep(5 * time.Second)
}