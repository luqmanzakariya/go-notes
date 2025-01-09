package belajargolanggoroutine

/* === Atomic ===
- Timer adalah representasi satu kejadian
- Ketika waktu timer sudah expire, maka event akan dikirim ke dalam channel
- Untuk membuat Timer kita bisa menggunakan time.NewTimer(duration)
*/

/* === time.After ===
- Kadang kita hanya butuh channel nya saja, tidak membutuhkan data Timernya
- Untuk melakukan hal itu kita bisa menggunakan function time.After(duration)
*/

/* === time.AfterFunc() ===
- Kadang ada kebutuhan kita ingin menjalankan function dengan delay waktu tertentu
- Kita bisa memanfaatkan Timer dengan menggunakan function time.AfterFunc()
- Kita tidak perlu lagi menggunakan channel nya, cukup kirim kan function yang akan dipanggil
	ketika Timer mengirim kejadiannya
*/

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T){
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <- timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T){
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <- channel
	fmt.Println(time)
}

func TestAfterFunction(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5 * time.Second, func () {
		fmt.Println("Print before grup done", time.Now())
		group.Done()
	})
	fmt.Println("Print outside Loop", time.Now())

	group.Wait()
}