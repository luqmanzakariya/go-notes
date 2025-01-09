package belajargolanggoroutine

/* === Atomic ===
- Go-Lang memiliki package yang bernama sync/atomic
- Atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada
	proses concurrent
- Contoh sebelumnya kita telah menggunakan Mutex untuk melakukan locking ketika ingin menaikkan
	angka di counter. Hal ini sebenarnya bisa digunakan menggunakan Atomic package
- Ada banyak sekali function di atomic package, kita bisa eksplore sendiri di halaman dokumentasinya
*/

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				/* 	
						Untuk memanipulasi tipe data yang merupakan primitive (string, integer, dll) kita
						menggunakan atomic dibandingkan dengan mutex (gak perlu capek capek).

						Kecuali jika tipe data merupakan struct atau lainnya lebih cocok untuk menggunakan
						mutex
				*/
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter: ", x)
}
