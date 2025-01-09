package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

/* === Buffered Channel ===
- Seperti yang dijelaskan sebelumnya, bahwa secara default channel itu hanya bisa menerima 1 data
- Artinya jika kita menambah data ke-2, maka kita akan diminta menunggu sampai data ke-1 ada yang
	mengambil
- Kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita
	menggunakan channel, maka otomatis pengirim akan ikut lambat juga
- Untungnya ada Buffered Channel, yaitu buffer yang bisa digunakan untuk menampung data antrian di
	Channel
*/

/* === Buffered Capacity ===
- Kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
- Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer
- Jika kita mengirim data ke 6, maka kita diminta untuk menunggu sampai ada buffer yang kosong
- Ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari yang mengirim data
*/

/* === Membuat Buffered Channel ===
channel := make(chan string, 3) // 3 disini adalah jumlah buffer channel tersebut
fmt.Println(cap(channel)) // Melihat panjang buffer
fmt.Println(len(channel)) // Melihat jumlah data di buffer
*/

func TestBufferedChannel(t *testing.T){
	channel := make(chan string, 3)
	defer close(channel)

	go func(){
		channel <- "Eko"
		channel <- "Kurniawan"
		channel <- "Khannedy"
	}()


	go func(){
		fmt.Println(<- channel)
		fmt.Println(<- channel)
		fmt.Println(<- channel)
	}()
	
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}