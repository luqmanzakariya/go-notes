package main

import (
	"fmt"
	go_say_hello "github.com/luqmanzakariya/go-modules-example/v2"
)

func main () {
	res := go_say_hello.SayHello("Luqman")
	fmt.Println(res)
}