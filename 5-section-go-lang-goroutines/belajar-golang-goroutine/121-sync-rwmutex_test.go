package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/* === RWMutex (Read Write Mutex)
- Kadang ada kasus dimana kita ingin melakukan locking tidak hanya pada proses mengubah data, tapi juga
	membaca data
- Kita sebenarnya bisa menggunakan Mutex saja, namun masalahnya nanti akan rebutan antara proses membaca
	dan mengubah
- Di Go-Lang telah disediakan struct RWMutex(Read Write Mutex) untuk menangani hal ini, dimana Mutex jenis
	ini memiliki dua lock, lock untuk Read dan lock untuk Write
*/

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int){
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T){
	account := BankAccount{}

	for i:=0; i<100; i++ {
		go func(){
			for j:=0; j<100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}