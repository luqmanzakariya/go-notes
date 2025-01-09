package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"time"
	"testing"
)

/* === Deadlock
- Hati-hati saat membuat aplikasi yang parallel atau concurrent, masalah yang sering kita hadapi
	adalah Deadlock
- Deadlock adalah keadaan dimana sebuah proses goroutine saling menunggu lock sehingga tidak ada satupun
	goroutine yang bisa jalan
- Sekarang kita coba simulasikan proses deadlock
*/

/* === Contoh Simulasi Deadlock (1) === */
type UserBalance struct {
	sync.Mutex
	Name	string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock(){
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int){
	user.Balance = user.Balance + amount
}

func Transfer (user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user 1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T){
	user1 := UserBalance{
		Name: "Eko",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name: "Budi",
		Balance: 1000000,
	}

	// Transfer(&user1, &user2, 100000) // Tanpa goroutine data expected
	
	// Data tidak expected
	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)
	
	time.Sleep(3 * time.Second)

	fmt.Println("User ", user1.Name, " Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, " Balance ", user2.Balance)
}