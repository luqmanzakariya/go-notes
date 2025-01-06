package main

import (
	"fmt"
	"reflect"
)

/* === Package reflect ===
- Dalam bahasa pemorgraman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode kita
  pada saat aplikasi sedang berjalan
- Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect
- Fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video, Anda bisa eksplorasi package
  reflect ini secara otodidak
- Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
*/

/* === Kode Program StructTag ===
- Didalam reflection terdapat fitur struct tag (tag didalam field)
- Cara menambahkan tag dengan `<nama tag>:<value tag>`
- Tag ini bisa dibaca dengan package reflection
- Fitur ini Sangat cocok untuk validation
*/

type Sample struct {
	Name string `required:"true" max:"10"` // Contoh structTag
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required")) // Get structTag
		fmt.Println(structField.Tag.Get("max"))      // Get structTag
	}
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}

	return result
}

// func main() {
// 	readField(Sample{"Eko"})
// 	readField(Person{"Eko", "Pamulang", "eko@mail.com"})

// 	person := Person {
// 		Name: "Eko",
// 		Address:  "Jakarta",
// 		Email: "",
// 	}

// 	fmt.Println("====== isValid ======")
// 	fmt.Println(isValid(person))
// }
