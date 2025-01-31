package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

/* === JSON Array ===
- Selain tipe dalam bentuk Object, biasanya dalam JSON, kita kadang menggunakan tipe data Array
- Array di JSON mirip dengan Array di Javascript, dia bisa berisikan tipe data primitif, atau tipe data kompleks
	(Object atau Array)
- Di Go-Lang, JSON Array direpresentasikan dalam bentuk slice
- Konversi dari JSON atau ke JSON dilakukan secara otomatis oleh package json menggunakan tipe data slice
*/

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type CustomerNew struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

func TestJSONArray(t *testing.T) {
	customer := CustomerNew{
		FirstName:  "Eko",
		MiddleName: "Kurniawan",
		LastName:   "Khannedy",
		Age:        30,
		Married:    true,
		Hobbies:    []string{"Gaming", "Reading", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"Kurniawan","LastName":"Khannedy","Age":30,"Married":true,"Hobbies":["Gaming","Reading","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &CustomerNew{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := CustomerNew{
		FirstName: "Eko",
		Addresses: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
			{
				Street:     "Jalan Lagi Dibangun",
				Country:    "Indonesia",
				PostalCode: "8888",
			},
		},
	}

	byte, _ := json.Marshal(customer)
	fmt.Println(string(byte))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Eko","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"8888"}]}`
	jsonBytes := []byte(jsonString)

	customer := &CustomerNew{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

/* === Decode JSON Array ===
- Selain menggunakan Array pada attribute di Object
- Kita juga bisa melakukan encode atau decode langsung JSON Array nya
- Encode dan Decode JSON Array bisa menggunakan tipe data slice
*/

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"8888"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJsonArrayComplex(t *testing.T) {
 	addresses := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "Jalan Lagi Dibangun",
			Country:    "Indonesia",
			PostalCode: "8888",
		},
	}

	byte, _ := json.Marshal(addresses)
	fmt.Println(string(byte))
}
