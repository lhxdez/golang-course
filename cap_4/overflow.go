package main

import (
	"fmt"
)

func main() {

	var i uint16
	i = 65535 // limite dos 16 bits
	fmt.Println(i)
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)
}
