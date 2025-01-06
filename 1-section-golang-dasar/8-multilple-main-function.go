/*
	Multiple Main Function
	
	1. Di Golang, function dalam module/project adalah unik,
	artinya kita tidak boleh membuat nama function yang sama
	2. Maka apabila kita membuat file baru, kemudian membuat function
	yang sama contoh main, maka kita tidak bisa melakukan build module.
	Karena function tersebut duplikat dengan yang ada di main function

*/

package main

import "fmt"

func multipleMain () {
	fmt.Print("yuhuu")
}