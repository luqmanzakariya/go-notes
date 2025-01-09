package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* === Mutext (Mutual Exclusion)
- Untuk mengatasi masalah race condition tersebut, di Go-Lang terdapat sebuah struct bernama sync.Mutex
- Mutex bisa digunakan untuk melakukan locking dan unlocking, dimana ketika kita melakukan locking terhadap
	mutex, maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan unlock
- Dengan demikian, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1 goroutine yang
	diperbolehkan, setelah goroutine tersebut melakukan unlock, baru goroutine selanjutnya diperbolehkan
	melakukan lock lagli
- Ini sangat cocok sebagai solusi ketika ada masalah race condition yang sebelumnya kita hadapi
*/

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock() // Untuk menlock sisa dari satu looping yang masuk kesini
				x = x + 1
				mutex.Unlock() // Setelah satu mendapatkan increment baru yang lain akan di unlock kembali
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter: ", x)
}
