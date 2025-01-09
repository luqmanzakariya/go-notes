package belajargolanggoroutine

import (
	"testing"
	"time"
	"fmt"
)


/* === Channel Sebagai Parameter ===
- Saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan menerima data
	dari channel tersebut
- Kadang kita ingin memberi tahu terhadap function, misal bahwa channel tersebut hanya digunakan untuk
	mengirim data, atau hanya dapat digunakan untuk menerima data
- Hal ini bisa kita lakukan di parameter dengan cara menandai apakah channel ini digunakan untuk in
	(mengirim data) atau out (menerima data)
*/

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Eko Kurniawan Khannedy"
}

func OnlyOut(channel <-chan string) {
	data := <- channel
	fmt.Println(data) 
}

func TestInOutChannel(t *testing.T){
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}