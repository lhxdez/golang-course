// - golang.org/ref/spec
// - Numa declaração de constantes, o identificador iota representa números sequenciais.
// - Na prática.
//     - iota, iota + 1, a = iota b c, reinicia em cada const, _

package main

import (
	"fmt"
)

const (
	a = iota + 10000000
	_
	c
	x
	_
	z
)

func iota() {
	fmt.Println(a, c, x, z)
}
