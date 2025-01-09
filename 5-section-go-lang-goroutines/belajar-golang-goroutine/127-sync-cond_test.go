package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* === Cond ===
- Cond adalah implementasi locking berbasis kondisi.
- Cond membutuhkan Locker (bisa menggunakan Mutex atau RWMutex) untuk implementasi locking nya,
	namun berbeda dengan Locker biasanya, di Cond terdapat function wait() untuk menunggu apakah
	perlu menunggu atau tidak
- Function Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu
	lagi, sedangkan function Broadcast() digunakan untuk memberi tahu semua goroutine agar tidak
	perlu menunggu lagi
- Untuk membuat Cond, kita bisa menggunakan function sync.NewCond(Locker)
*/

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int){
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	
	// code setelah wait ini tidak akan berjalan apabila tidak terdapat signal
	fmt.Println("Done", value)

	cond.L.Unlock()
}

func TestCond(t *testing.T){
	for i:= 0; i<10; i++ {
		go WaitCondition(i)
	}

	// Implementasi signal, akan ngetrigger signal satu persatu
	go func(){
		for i:= 0; i<10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// Implementasi broadcast, akan ngetrigger semua signal dalam 1 waktu
	// go func(){
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
}