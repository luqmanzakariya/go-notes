package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

/* === Context With Cancel ===
- Selain menambahkan value ke context, kita juga bisa menambahkan sinyal cancel ke context
- Kapan sinyal cancel diperlukan dalam context?
- Biasanya ketika kita butuh menjalankan proses lain, dan kita ingin memberi sinyal cancel ke proses tersebut
- Biasanya proses ini berupa goroutine yang berbeda, sehingga dengan mudah jika kita ingin membatalkan eksekusi
	goroutine, kita bisa mengirim sinyal cancel ke context nya
- Namun ingat, goroutine yang menggunakan context tetap harus melakukan pengecekan terhadap context nya,
	jika tidak, tidak ada gunanya
- Untuk membuat context dengan cancel signal, kita bisa menggunakan function
	context.WithCancel(parent)
*/

// Contoh Goroutine Leak (1)
func CreateCounter() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

// Contoh Testing Goroutine Leak (1)
func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	destination := CreateCounter()
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
	/*
		printnya menjadi 3 (bertambah satu) karena terdapat satu function yang tidak tahu kapan
		harus berhentinya. Hal ini sangat berbahaya apabila tidak di handle. Sebagai contoh apabila terdapat
		1000 request maka goroutine akan terus bertambah. Hal ini dapat menyebabkan konsumsi memori yang
		berlebihan dan kondisi aplikasi crash yang agak sulit untuk di trace mengapa aplikasi tersebut crash
	*/
}

func CreateCounterWithoutLeak(ctx context.Context) chan int {
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
			}
		}
	}()

	return destination
}

func TestContextWithCancelWithoutLeak(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounterWithoutLeak(ctx)

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel() // mengirim sinyal cancel ke context
	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
	// Go routine sudah kembali balik menjadi 2
}