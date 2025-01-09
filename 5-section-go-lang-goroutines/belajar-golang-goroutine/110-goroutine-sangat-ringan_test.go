package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

/* === Goroutine Sangat Ringan ===
- Seperti yang sebelumnya dijelaskan, bahwa goroutine itu sangat ringan
- Kita bisa membuat ribuan, bahkan sampai jutaan goroutine tanpa takut boros memory
- Tidak seperti thread yang ukurannya berat, goroutine sangatlah ringan
*/

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		// DisplayNumber(i) // Tanpa goroutine
		go DisplayNumber(i) // Dengan goroutine
	}

	time.Sleep(5 * time.Second)
}

/* 	
	# Pertanyaan:	kenapa displaynya tidak berurutan dengan menggunakan go routine?
	# Jawaban		: Karena running dengan go routine, dia runningnya asynchronous, walaupun sebenarnya
								dia itu concurrent (harusnya berurutan 1, 2, 3, ..., dst) tapi karena laptop saya
								multicore, ada sekitar 10 processor, jadi dia gak running secara concurrent tapi juga
								secara parallel. Oleh karena itu angkanya bisa berubah ubah jadi tidak terstruktur.
*/

/*
	# Pertanyaan: Kenapa goroutine disebut sangat ringan?
	# Jawaban		:	Berdasarkan diatas test looping yang dilakukan sebanyak 100,000 kali program tersebut
								selesai tanpa terdapat error. Hal ini yang membuat goroutine go lang sangat ringan,
								berbeda dengan bahasa pemograman seperti java yang apabila dilakukan sekitar 50rb
								mungkin sudah akan muncul error
*/