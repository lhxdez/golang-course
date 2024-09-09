// - Crie uma variável de tipo string utilizando uma raw string literal.
// - Demonstre-a.

package main

import (
	"fmt"
)

func ex5() {
	x := `isto
	é
		uma coisa
			muito doida`
	fmt.Println(x)
}
