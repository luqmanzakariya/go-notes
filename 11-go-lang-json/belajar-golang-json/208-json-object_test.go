package main

import (
	// "encoding/json"
	// "fmt"
	"encoding/json"
	"fmt"
	"testing"
)

/* === JSON Object ===
- Pada materi sebelumnya kita melakukan encode data seperti string, number, boolean, dan tipe data primitif lainnya
- Walaupun memang bisa dilakukan, karena sesuai dengan kontrak interface{}, namun tidak sesuai dengan kontrak JSON
- Jika mengikuti kontrak json.org, data JSON bentuknya adalah Object dan Array
- Sedangkan value nya baru berupa
*/

/* === Struct ===
- JSON Object di Go-Lang direpresentasikan dengan tipe data Struct
- Dimana tiap attribute di JSON Object merupakan attribute di Struct
*/

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJSONObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
		Age:        30,
		Married:    true,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
