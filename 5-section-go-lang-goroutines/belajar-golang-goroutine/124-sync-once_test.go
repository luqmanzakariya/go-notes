package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
)

/* === Once ===
- Once adalah fitur di Go-Lang yang bisa kita gunakan untuk memastikan bahwa sebuah function
	di eksekusi hanya sekali
- Jadi berapa banyak pun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama
	yang bisa mengeksekusi function nya
- Goroutine yang lain akan dihiraukan, artinya function tidak akan dieksekusi lagi
*/

/* === Contoh Once === */
var counter = 0

func OnlyOnce() {
	counter++
}


func TestOnce(t *testing.T){
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i:= 0; i<100; i++ {
		go func() {
			group.Add(1) 
			// karena ada once.do maka counter hanya dieksekusi cuma 1 kali saja
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("counter", counter)
}