package main

import (
	"fmt"
)

func main() {
	s := "ascii éøâ 香"

	// caracter por caracter
	for _, v := range s {
		fmt.Printf("%b - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	// byte por byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}

}
