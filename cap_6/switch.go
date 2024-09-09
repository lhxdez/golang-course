// - Switch:
//     - pode avaliar uma expressão
//         - switch statement == case (value)
//         - default switch statement == true (bool)
//     - não há fall-through por padrão
//     - criando fall-through
//	   		-
//     - default
//     - cases compostos

// - Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.

package main

import (
	"fmt"
)

func switchh() {
	qm := "joana"

	switch qm {
	case "joana":
		fmt.Println("joana esta presente")
		fallthrough
	case "marcos":
		fmt.Println("marcos esta presente")
	case "adonis":
		fmt.Println("adonis esta presente")
	}

}
