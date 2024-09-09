package main

import (
	"fmt"
)

func ascii() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("%d - %v\n", x, string(x))
	}
}
