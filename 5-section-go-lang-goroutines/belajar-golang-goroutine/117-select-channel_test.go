package belajargolanggoroutine

import (
	"fmt"
	"testing"
)

/* === Select Channel ===
- Kadang ada kasus dimana kita membuat beberapa channel, dan menjalankan beberapa goroutine
- Lalu kita ingin mendapatkan data dari semua channel tersebut
- Untuk melakukan hal tersebut, kita bisa menggunakan select channel di Go-Lang
- Dengan select channel, kita bisa memilih data tercepat dari beberapa channel, jika data
	datang secara bersamaan di beberapa channel, maka akan dipilih secara random
*/

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}

		/*
			Agar tidak terjadi dead lock maka perlu ditambahkan counter. Karena jumlah channel hanya terdapat 2,
			jadi akan terjadi deadlock pada saat looping ketiga
		*/
		if counter == 2 {
			break
		}
	}
}
