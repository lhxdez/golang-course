package main

import (
	"fmt"
)

func break_e_continue() {
	for i := 0; i < 20; i++ {
		if i == 3 {
			// faz isso quando o número não é par
			continue
		}
		fmt.Println(i)
	}
}

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	x := 0

// 	for x <= 20 {

// 		x++

// 		if (x % 2 != 0) {
// 			continue
// 		}

// 		fmt.Println(x)

// 	}

// }
