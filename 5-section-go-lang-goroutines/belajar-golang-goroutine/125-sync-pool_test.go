package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* === Pool ===
- Pool adalah implementasi design pattern bernama object pool pattern
- Sederhananya, design pattern Pool ini digunakan untuk menyimpan data, selanjutnya untuk
	menggunakan datanya, kita bisa mengambil dari pool, dan setelah selesai menggunakan datanya,
	kita bisa menyimpan kembali ke Pool nya
- Implementasi Pool di Go-Lang ini sudah aman dari problem race condition
*/

/* === Membuat Data Pool Otomatis === 
	var pool = sync.Pool {
		New: func()interface{}{
			return "New"
		}
	}

*/


func TestPool(t *testing.T) {
	pool := sync.Pool{}

	pool.Put("Eko")
	pool.Put("Kurniawan")
	pool.Put("Khannedy")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
}
