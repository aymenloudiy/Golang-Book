package main

import (
	"example/greeting"
	"fmt"
)

func main() {
	message:=greeting.Hello("nemyA")
	fmt.Println(message)
}