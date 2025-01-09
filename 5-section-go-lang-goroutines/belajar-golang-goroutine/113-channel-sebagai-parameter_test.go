package belajargolanggoroutine

import (
	"testing"
	"time"
	"fmt"
)

/* === Channel Sebagai Parameter ===
- Dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter
- Sebelumnya kita tahu bahkan di Go-Lang by default parameter adalah pass by value, artinya value akan diduplikasi
	lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kita bisa gunakan pointer
	(agar pass by reference)
- Berbeda dengan Channel, kita tidak perlu melakukan hal tersebut
*/

func GiveMeResponse(channel chan string){
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan Khannedy"
}

func TestChannelAsParameter(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data) 

	time.Sleep(5 * time.Second)
}