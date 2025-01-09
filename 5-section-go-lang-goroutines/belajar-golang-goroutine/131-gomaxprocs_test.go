package belajargolanggoroutine

/* === GOMAXPROCS ===
- Sebelumnya diawal kita sudah bahas bahwa goroutine itu sebenarnya dijalankan di dalam Thread
- Pertanyaannya, seberapa banyak Thread yang ada di Go-Lang ketika aplikasi kita berjalan?
- Untuk mengetahui berapa jumlah Tread, kita bisa menggunakan GOMAXPROCS, yaitu sebuah function
	di package runtime yang bisa kita gunakan untuk mengubah jumlah thread atau mengambil jumlah
	thread
- Secara default, jumlah thread di Go-Lang itu sebanyak jumlah CPU di komputer kita
- Kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function runtime.NumCpu()
*/

/* === Mengubah Jumlah Thread ===
- Go routine diubah berdasarkan jumlah go routine
- CPU tidak bisa diubah kecuali tambah CPU hardware
- Untuk mengubah jumlah thread kita dapat menggunakan (vertical scaling):
	runtime.GOMAXPROCS(num int)
*/

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T){
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu) // total cpu

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread) // by default total thread = total cpu
	
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine) /* minimal ada 2 go routine yang berjalan, 1 untuk
																										running code, 1 lagi untuk garbage collection 
																										ketika menghapus data di memory sama si go
																										run time
																									*/
	group := sync.WaitGroup{}
	for i := 0; i<100; i++ {
		group.Add(i)

		go func(){
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalGoroutine = runtime.NumGoroutine()
	fmt.Println("Total Goroutine setelah looping go routine", totalGoroutine) // go routine + 100 (dari loop)
	group.Wait()
}

func TestChangeThreadNumber(t *testing.T){
	runtime.GOMAXPROCS(20) // Mengubah jumlah thread menjadi 20
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)
}