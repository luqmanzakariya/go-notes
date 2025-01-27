package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/* === Context With Timeout ===
- Selain menambahkan value ke context, dan juga sinyal cancel, kita juga bisa menambahkan sinyal cancel
	ke context secara otomatis dengan menggunakan pengaturan timeout
- Dengan menggunakan pengaturan timeout, kita tidak perlu melakukan eksekusi cancel secara manual,
	cancel akan otomatis di eksekusi jika waktu timeout sudah terlewati
- Penggunaan context dengan timeout sangat cocok ketika misal kita melakukan query ke database atau
	http api, namun ingin menentukan batas maksimal timeout nya
- Untuk membuat context dengan cancel signal secara otomatis menggunakan timeout, kita bisa menggunakan
	function:
		context.WithTimeout(parent, duration)
*/

// Context With Timeout (1)

func CreateCounterWithTimeout(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulasi Slow
			}
		}
	}()

	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5 * time.Second)
	defer cancel() // Function ini kalo berfungsi kalo function lebih cepet akan membatalkan context

	destination := CreateCounterWithTimeout(ctx)
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}