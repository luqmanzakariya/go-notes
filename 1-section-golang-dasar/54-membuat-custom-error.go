package main

/* === Membuat Custom Error ===
- Karena error adalah sebuah interface, jadi jika kita ingin membuat error sendiri, kita bisa membuat struct
  yang mengikuti kontrak dari interface error
*/

// import "fmt"

// "errors"
// "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "eko" {
		return &notFoundError{"data not found"}
	}

	return nil
}

// func main() {
// 	err := SaveData("eko", nil)

// 	if err != nil {
// 		// terjadi error
// 		if validationError, ok := err.(*validationError); ok {
// 			fmt.Println("validation error: ", validationError.Error())
// 		} else if notFoundErr, ok := err.(*notFoundError); ok {
// 			fmt.Println("not found error: ", notFoundErr.Error())
// 		} else {
// 			fmt.Println("unknown error", err.Error())
// 		}
// 	} else {
// 		fmt.Println("Success")
// 	}
// }
