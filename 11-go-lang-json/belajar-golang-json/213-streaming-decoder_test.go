package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

/* === Streaming Decoder ===
- Sebelumnya kita belajar package json dengan melakukan konversi data JSON yang sudah dalam bentuk
	variable dan data string atau []byte
- Pada kenyataannya, kadang data JSON nya berasar dari Input berupa
	io.Reader (File, Network, Request Body)
- Kita bisa saja membaca semua datanya terlebih dahulu, lalu disimpan di variable, baru lakukan
	konversi dari JSON, namun hal ini sebenarnya tidak perlu dilakukan, karena package json memiliki fitur
	untuk membaca data dari Stream
*/

/* === json.Decoder ===
- Untuk membuat json Decoder, kita bisa menggunakan function json.NewDecoder(reader)
- Selanjutnya untuk membaca isi input reader dan konversikan secara langsung ke data di Go-Lang, cukup
	gunakan function Decode(interface{})
*/

func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("Customer.json")
	decoder := json.NewDecoder(reader)

	customer := &CustomerNew{}
	decoder.Decode(customer)

	fmt.Println(customer)
}
