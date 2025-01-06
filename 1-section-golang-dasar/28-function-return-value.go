package main

// import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Tidak Punya nama"
	} else {
		return "Halo" + " " + name
	}
}

// func main() {
// 	result := getHello("")
// 	fmt.Println(result)
// }
