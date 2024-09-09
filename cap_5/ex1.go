// - Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

package main

import (
	"fmt"
)

func ex1() {
	x := 100
	fmt.Printf("%d, %#x, %b", x, x, x)
}
