package belajargolanggoroutine

/* === time.Ticker (Mirip SetInterval di javascript) ===
- Ticker adalah representasi kejadian yang berulang
- Ketika waktu ticker sudah expire, maka event akan dikirim ke dalam channel
- Untuk membuat ticker, kita bisa menggunakan time.NewTicker(duration)
- Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop()
*/

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T){
	ticker := time.NewTicker(1 * time.Second)

	go func (){
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}


/* === time.Tick() ===
- Kadang kita tidak membutuhkan data Ticker nya, kita hanya butuh channel nya saja
- Jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan
	mengembalikan Ticker, hanya mengembalikan channel timer nya saja
*/

func TestTick(t *testing.T){
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}