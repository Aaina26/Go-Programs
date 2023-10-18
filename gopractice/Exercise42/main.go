package main

import (
	"fmt"
)

func main() {
	ar := [5]int{22, 33, 44, 55, 66}

	for _, val := range ar {
		fmt.Printf("Value is: %d   Type is: %T\n", val, val)
	}
}
