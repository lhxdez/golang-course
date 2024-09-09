// - Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
// - Demonstre estes valores.

package main

import (
	"fmt"
)

const (
	_ = 1994 + iota
	b
	c
	d
	e
)

func ex6() {
	fmt.Println(b, c, d, e)
}
