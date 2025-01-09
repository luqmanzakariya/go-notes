package belajargolanggoroutine

/* === Map ===
- Go lang memiliki sebuah struct bernama sync.Map
- Map ini mirip Go-Lang map, namun yang membedakannya, Map ini aman untuk menggunakan concurrent
	menggunakan goroutine
- Ada beberapa function yang bisa kita gunakan di Map:
	- Store(key, value)						: untuk menyimpan data ke Map
	- Load(key) 									: Untuk mengambil data dari Map menggunakan key
	- Delete(key)									: Untuk menghapus data di Map menggunakan key
	- Range(function(key, value)) : Digunakan untuk melakukan iterasi seluruh data di Map
*/

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range((func(key, value interface{}) bool {
		fmt.Println(key, " : ", value)
		return true
	}))
}
