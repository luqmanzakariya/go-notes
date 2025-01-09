package belajargolanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
)

/* === Buffered Channel ===
- Kadang-kadang ada kasus channel dikirim data secara terus menerus oleh pengirim
- Dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
- Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan range ketika menerima data
	dari channel
- Ketika sebuah channel di close(), maka secara otomatis perulangan tersebut akan berhenti
- Ini lebih sederhana dari pada kita melakukan pengecek channel secara manual
*/

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perlulangan ke " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}
